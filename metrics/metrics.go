package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func help(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Fprintln(w, "Metrics Help")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "	POST /v1/metrics/node/{nodename}/")
	fmt.Fprintln(w, "	POST /v1/metrics/nodes/{nodename}/process/{processname}/")
	fmt.Fprintln(w, "	GET /v1/analytics/nodes/average")
	fmt.Fprintln(w, "	GET /v1/analytics/processes/")
	fmt.Fprintln(w, "	GET /v1/analytics/processes/{processname}/")
	fmt.Fprintln(w, "")
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	log.Println(r.Form) // print form information in server side
	log.Println("path", r.URL.Path)
	log.Println("scheme", r.URL.Scheme)
	log.Println(r.Form["url_long"])
	for k, v := range r.Form {
		log.Println("key:", k)
		log.Println("val:", strings.Join(v, ""))
	}

	// response
	fmt.Fprintf(w, "Hello")
}

func main() {

	PORT := ":9911"

	log.Println("Listening on ", PORT)

	http.HandleFunc("/", help)
	http.HandleFunc("/v1/metrics/node", help)
	http.HandleFunc("/v1/metrics/nodes", help)
	http.HandleFunc("/v1/analytics/nodes", help)
	http.HandleFunc("/v1/analytics/processes", help)

	err := http.ListenAndServe(PORT, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
