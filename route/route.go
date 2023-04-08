package route

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/sraynitjsr/entity"
	"github.com/sraynitjsr/repository"
)

var repo repository.PostRepository = repository.NewPostRepository()

func Home(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte(`{"Welcome Home"}`))
}

func GetPosts(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-type", "application/json")

	posts, err := repo.FindAll()

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(`{"error":"error getting the posts"}`))
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(posts)
}

func AddPosts(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-type", "application/json")
	var newPost entity.Post
	err := json.NewDecoder(request.Body).Decode(&newPost)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(`{"error":"error unmarshalling the request"}`))
		return
	}
	newPost.ID = int64(rand.Int())
	repo.Save(&newPost)
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(newPost)
}
