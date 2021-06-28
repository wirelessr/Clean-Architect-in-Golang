package controller

import (
	"encoding/json"
	"net/http"

	"example.com/clean-arch/entity"
	"example.com/clean-arch/errors"
	"example.com/clean-arch/service"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	err = postService.Validate(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	result, err := postService.Create(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
