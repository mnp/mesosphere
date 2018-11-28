// Partially lifted from circleci's go demo.
// Starts a web server and db connection.

package main

import (
	"log"
	"net/http"
	"github.com/mnp/mesosphere/service"
)

func main() {
	HOSTPORT := ":9911"

	log.Println("Listening on ", HOSTPORT)

	//TODO	db := SetupDB()

	server := service.NewServer( /* db */ )
	http.HandleFunc("/", server.ServeHTTP)
	err := http.ListenAndServe(HOSTPORT, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
