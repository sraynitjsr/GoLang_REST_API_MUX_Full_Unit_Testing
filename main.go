package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sraynitjsr/route"
)

func main() {
	fmt.Println("GoLang Gorilla MUX REST API")

	router := mux.NewRouter()

	router.HandleFunc("/", route.Home)

	router.HandleFunc("/posts", route.GetPosts).Methods("GET")

	router.HandleFunc("/posts", route.AddPosts).Methods("POST")

	log.Println("Starting Server at Port 8000")
	log.Fatalln(http.ListenAndServe(":8000", router))
}
