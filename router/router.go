package router

import "net/http"

type Router interface {
	Home(uri string, f func(responseWriter http.ResponseWriter, request *http.Request))
	Get(uri string, f func(responseWriter http.ResponseWriter, request *http.Request))
	Post(uri string, f func(responseWriter http.ResponseWriter, request *http.Request))
	Start(port string)
}
