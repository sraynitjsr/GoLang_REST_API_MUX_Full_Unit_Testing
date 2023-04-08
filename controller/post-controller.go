package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sraynitjsr/entity"
	"github.com/sraynitjsr/errors"
	"github.com/sraynitjsr/service"
)

var svc service.PostService = service.NewPostService()

type PostController interface {
	Home(responseWriter http.ResponseWriter, request *http.Request)
	GetPosts(responseWriter http.ResponseWriter, request *http.Request)
	AddPosts(responseWriter http.ResponseWriter, request *http.Request)
}

type PostControllerImpl struct{}

func NewPostController() PostController {
	return &PostControllerImpl{}
}

func (pc *PostControllerImpl) Home(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte(`{"Welcome Home"}`))
}

func (pc *PostControllerImpl) GetPosts(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-type", "application/json")

	posts, err := svc.FindAll()

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(responseWriter).Encode(errors.ServiceError{Message: "error in getting posts"})
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(posts)
}

func (pc *PostControllerImpl) AddPosts(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-type", "application/json")

	var newPost entity.Post

	err := json.NewDecoder(request.Body).Decode(&newPost)

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(responseWriter).Encode(errors.ServiceError{Message: "error unmarshalling the request"})
		return
	}

	validationError := svc.Validate(&newPost)

	if validationError != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(responseWriter).Encode(errors.ServiceError{Message: validationError.Error()})
		return
	}

	result, createError := svc.Create(&newPost)
	if createError != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(responseWriter).Encode(errors.ServiceError{Message: "error saving the post"})
		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	json.NewEncoder(responseWriter).Encode(result)
}
