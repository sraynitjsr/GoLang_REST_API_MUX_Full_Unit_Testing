package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("GoLang Gorilla MUX REST API")

	router := mux.NewRouter()

	router.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(responseWriter, "Home Handler, Welcome...............")
	})
	log.Println("Starting Server at Port 8000")
	log.Fatalln(http.ListenAndServe(":8000", router))
}
