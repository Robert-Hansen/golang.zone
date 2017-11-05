package controllers

import (
	"github.com/steffen25/golang.zone/app"
	"github.com/steffen25/golang.zone/repositories"
	"net/http"
)

type TestController struct {
	*app.App
	repo repositories.TestPostRepository
}

func NewTestController(a *app.App, test repositories.TestPostRepository) *TestController {
	return &TestController{a, test}
}

func (test *TestController) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := test.repo.GetAllPosts()

	if err != nil {
		NewAPIError(&APIError{false, "Could not fetch posts", http.StatusBadRequest}, w)
		return
	}

	if len(posts) == 0 {
		NewAPIResponse(&APIResponse{Success: false, Message: "Could not find posts", Data: posts}, w, http.StatusNotFound)
		return
	}

	NewAPIResponse(&APIResponse{Success: true, Data: posts}, w, http.StatusOK)
}
