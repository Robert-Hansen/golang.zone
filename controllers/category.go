package controllers

import (
	"net/http"
	"time"

	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/steffen25/golang.zone/app"
	"github.com/steffen25/golang.zone/models"
	"github.com/steffen25/golang.zone/repositories"
	"github.com/steffen25/golang.zone/util"
)

type CategoryController struct {
	*app.App
	repositories.CategoryRepository
}

func NewCategoryController(a *app.App, cr repositories.CategoryRepository) *CategoryController {
	return &CategoryController{a, cr}
}

func (cc *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := cc.CategoryRepository.GetAll()
	if err != nil {
		NewAPIError(&APIError{false, "Could not fetch categories", http.StatusBadRequest}, w)
		return
	}

	if len(categories) == 0 {
		NewAPIResponse(&APIResponse{Success: false, Message: "Could not find categories", Data: categories}, w, http.StatusNotFound)
		return
	}

	NewAPIResponse(&APIResponse{Success: true, Data: categories}, w, http.StatusOK)
}

func (cc *CategoryController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		NewAPIError(&APIError{false, "Invalid request", http.StatusBadRequest}, w)
		return
	}
	category, err := cc.CategoryRepository.FindById(id)
	if err != nil {
		NewAPIError(&APIError{false, "Could not find category", http.StatusNotFound}, w)
		return
	}

	NewAPIResponse(&APIResponse{Success: true, Data: category}, w, http.StatusOK)
}

func (cc *CategoryController) GetByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	category, err := cc.CategoryRepository.FindByName(name)
	if err != nil {
		NewAPIError(&APIError{false, "Could not find category", http.StatusNotFound}, w)
		return
	}

	NewAPIResponse(&APIResponse{Success: true, Data: category}, w, http.StatusOK)
}

func (cc *CategoryController) Create(w http.ResponseWriter, r *http.Request) {

	j, err := GetJSON(r.Body)
	if err != nil {
		NewAPIError(&APIError{false, "Invalid request", http.StatusBadRequest}, w)
		return
	}

	name, err := j.GetString("name")
	if err != nil {
		NewAPIError(&APIError{false, "Name is required", http.StatusBadRequest}, w)
		return
	}

	name = util.CleanZalgoText(name)

	if len(name) < 3 {
		NewAPIError(&APIError{false, "Name is too short", http.StatusBadRequest}, w)
		return
	}

	category := &models.Category{
		Name:      name,
		CreatedAt: time.Now(),
	}

	err = cc.CategoryRepository.Create(category)
	if err != nil {
		NewAPIError(&APIError{false, "Could not create category", http.StatusBadRequest}, w)
		return
	}

	defer r.Body.Close()
	NewAPIResponse(&APIResponse{Success: true, Message: "Category created", Data: category}, w, http.StatusOK)
}

func (cc *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId, err := strconv.Atoi(vars["id"])
	if err != nil {
		NewAPIError(&APIError{false, "Invalid request", http.StatusBadRequest}, w)
		return
	}
	category, err := cc.CategoryRepository.FindById(categoryId)
	if err != nil {
		NewAPIError(&APIError{false, "Could not find category", http.StatusNotFound}, w)
		return
	}

	j, err := GetJSON(r.Body)
	if err != nil {
		NewAPIError(&APIError{false, "Invalid request", http.StatusBadRequest}, w)
		return
	}

	name, err := j.GetString("name")
	if err != nil {
		NewAPIError(&APIError{false, "Name is required", http.StatusBadRequest}, w)
		return
	}

	name = util.CleanZalgoText(name)

	if len(name) < 3 {
		NewAPIError(&APIError{false, "Name is too short", http.StatusBadRequest}, w)
		return
	}

	category.UpdatedAt = mysql.NullTime{Time: time.Now(), Valid: true}
	category.Name = name

	err = cc.CategoryRepository.Update(category)
	if err != nil {
		NewAPIError(&APIError{false, "Could not update category", http.StatusBadRequest}, w)
		return
	}

	NewAPIResponse(&APIResponse{Success: true, Message: "Category updated", Data: category}, w, http.StatusOK)
}
