package router

import "net/http"

type Router interface {
	HOME(uri string, f func(responseWriter http.ResponseWriter, request *http.Request))
	GET(uri string, f func(responseWriter http.ResponseWriter, request *http.Request))
	POST(uri string, f func(responseWriter http.ResponseWriter, request *http.Request))
	SERVE(port string)
}
