package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sraynitjsr/controller"
)

var (
	gorillaMuxRouter = mux.NewRouter()
	newController    = controller.NewPostController()
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (mr *muxRouter) Home(uri string, f func(responseWriter http.ResponseWriter, request *http.Request)) {
	gorillaMuxRouter.HandleFunc(uri, newController.Home)
}

func (mr *muxRouter) Get(uri string, f func(responseWriter http.ResponseWriter, request *http.Request)) {
	gorillaMuxRouter.HandleFunc(uri, newController.GetPosts).Methods("GET")
}

func (mr *muxRouter) Post(uri string, f func(responseWriter http.ResponseWriter, request *http.Request)) {
	gorillaMuxRouter.HandleFunc(uri, newController.AddPosts).Methods("POST")
}

func (mr *muxRouter) Start(port string) {
	log.Println("Starting Server at Port", port)
	log.Fatalln(http.ListenAndServe(port, gorillaMuxRouter))
}
