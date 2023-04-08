package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	gorillaMuxRouter = mux.NewRouter()
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (mr *muxRouter) HOME(uri string, f func(responseWriter http.ResponseWriter, request *http.Request)) {
	gorillaMuxRouter.HandleFunc(uri, f).Methods("GET")
}

func (mr *muxRouter) GET(uri string, f func(responseWriter http.ResponseWriter, request *http.Request)) {
	gorillaMuxRouter.HandleFunc(uri, f).Methods("GET")
}

func (mr *muxRouter) POST(uri string, f func(responseWriter http.ResponseWriter, request *http.Request)) {
	gorillaMuxRouter.HandleFunc(uri, f).Methods("POST")
}

func (mr *muxRouter) SERVE(port string) {
	log.Println("Starting Server at Port", port)
	log.Fatalln(http.ListenAndServe(port, gorillaMuxRouter))
}
