package main

import (
	"fmt"

	"github.com/sraynitjsr/controller"
	"github.com/sraynitjsr/router"
)

func main() {
	fmt.Println("GoLang Gorilla MUX REST API")

	myMuxRouter := router.NewMuxRouter()
	myController := controller.NewPostController()

	myMuxRouter.Get("/", myController.Home)
	myMuxRouter.Get("/posts", myController.GetPosts)
	myMuxRouter.Post("/posts", myController.AddPosts)

	myMuxRouter.Start(":8080")
}
