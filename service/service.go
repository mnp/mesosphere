package service

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

// ErrorResponse is returned by our service when an error occurs.
type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// Server contains all that is needed to respond to incoming requests, like a database. Other services like a mail,
// redis, or S3 server could also be added.
type Server struct {
	router *httprouter.Router
	//TOOD	db     *Database
}

// The ServerError type allows errors to provide an appropriate HTTP status code and message. The Server checks for
// this interface when recovering from a panic inside a handler.
type ServerError interface {
	HttpStatusCode() int
	HttpStatusMessage() string
}

// Metrics intake for one node
type NodeMetrics struct {
	Timeslice float32 `json:"timeslice"` // number of seconds this measurement represents
	Cpu       float32 `json:"cpu"`       // percentage used
	Mem       float32 `json:"mem"`       // percentage used
}

// Node cache: map of nodename -> metrics
// TODO: sync to DB
var nodes map[string]NodeMetrics

// NewServer initializes the service with the given Database, and sets up appropriate routes.
func NewServer( /* db *Database */ ) *Server {
	router := httprouter.New()
	server := &Server{
		router: router,
		//TODO	db:     db,
	}

	// initialize cache
	nodes = make(map[string]NodeMetrics)

	server.setupRoutes()
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) setupRoutes() {

	log.Printf("setting up %#v", s.router)

	s.router.GET("/", s.Help)
	s.router.POST("/v1/metrics/node/:nodename/", s.PostNode)
	s.router.GET("/v1/analytics/nodes/average/:timeslice", s.GetNodes)

	// TODO: Not implemented yet. Route to help.
	s.router.GET("/v1/analytics/processes/", s.Help)
	s.router.GET("/v1/analytics/processes/:processname/", s.Help)
	s.router.POST("/v1/metrics/nodes/:nodename/process/:processname/", s.Help)

	// By default the router will handle errors. But the service should always return JSON if possible, so these
	// custom handlers are added.

	s.router.NotFound = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("404: %#v", r)
			writeJSONError(w, http.StatusNotFound, "")
		},
	)

	s.router.HandleMethodNotAllowed = true
	s.router.MethodNotAllowed = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("401: %#v", r)
			writeJSONError(w, http.StatusMethodNotAllowed, "")
		},
	)

	s.router.PanicHandler = func(w http.ResponseWriter, r *http.Request, e interface{}) {
		serverError, ok := e.(ServerError)
		if ok {
			writeJSONError(w, serverError.HttpStatusCode(), serverError.HttpStatusMessage())
		} else {
			log.Printf("Panic during request: %v", e)
			writeJSONError(w, http.StatusInternalServerError, "")
		}
	}
}

func (s *Server) PostNode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var node NodeMetrics
	nodename := ps.ByName("nodename")

	log.Printf("PostNode: %#v / %#v", r, ps)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&node); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Error decoding JSON")
		return
	}

	log.Printf("PostNode: %#v", node)
	nodes[nodename] = node
}

func (s *Server) GetNodes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	averages := NodeMetrics{0.0, 0.0, 0.0}

	timeslice, err := strconv.ParseFloat(ps.ByName("timeslice"), 32)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Error decoding timeslice parameter - wanted a float")
		return
	}

	for _, v := range nodes {
		// TODO: filter node data by timeslice
		averages.Cpu += v.Cpu
		averages.Mem += v.Mem
	}

	averages.Timeslice = float32(timeslice)
	averages.Cpu /= float32(len(nodes))
	averages.Mem /= float32(len(nodes))

	writeJSON(w, http.StatusOK, &averages)
}

func (s *Server) Help(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm() // parse arguments, you have to call this by yourself
	fmt.Fprintln(w, "Metrics Help")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "	POST /v1/metrics/node/{nodename}/")
	fmt.Fprintln(w, "	POST /v1/metrics/nodes/{nodename}/process/{processname}/")
	fmt.Fprintln(w, "	GET /v1/analytics/nodes/average")
	fmt.Fprintln(w, "	GET /v1/analytics/processes/")
	fmt.Fprintln(w, "	GET /v1/analytics/processes/{processname}/")
	fmt.Fprintln(w, "")
}

// ===== JSON HELPERS ==================================================================================================

func writeJSON(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func writeJSONError(w http.ResponseWriter, statusCode int, message string) {
	if message == "" {
		message = http.StatusText(statusCode)
	}

	writeJSON(
		w,
		statusCode,
		&ErrorResponse{
			StatusCode: statusCode,
			Message:    message,
		},
	)
}

func writeJSONNotFound(w http.ResponseWriter) {
	writeJSONError(w, http.StatusNotFound, "")
}

func writeUnexpectedError(w http.ResponseWriter, err error) {
	writeJSONError(w, http.StatusInternalServerError, err.Error())
}
