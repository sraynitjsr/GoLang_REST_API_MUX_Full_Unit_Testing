package main

import (
	"fmt"

	"github.com/sraynitjsr/controller"
	"github.com/sraynitjsr/repository"
	"github.com/sraynitjsr/router"
	"github.com/sraynitjsr/service"
)

var (
	repo         = repository.NewFireStoreRepository()
	svc          = service.NewPostService(repo)
	myController = controller.NewPostController(svc)

	myMuxRouter = router.NewMuxRouter()
)

func main() {
	fmt.Println("GoLang Gorilla MUX REST API")

	myMuxRouter.GET("/", myController.Home)
	myMuxRouter.GET("/posts", myController.GetPosts)
	myMuxRouter.POST("/posts", myController.AddPosts)

	myMuxRouter.SERVE(":8080")
}
