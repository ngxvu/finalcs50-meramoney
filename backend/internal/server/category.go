package server

import (
	"encoding/json"
	"meramoney/backend/infrastructure/domains"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateCategory creates a new category
func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {

	var categoryRequest CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var category domains.Category
	category.Name = categoryRequest.Name
	category.Description = categoryRequest.Description

	if err := s.DB.Create(&category).Error; err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

// GetCategory retrieves a category by ID
func (s *Server) GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category domains.Category
	if err := s.DB.First(&category, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(category)
}

// GetAllCategories retrieves all categories
func (s *Server) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for pagination
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Retrieve categories with pagination
	var categories []domains.Category
	if err := s.DB.Limit(pageSize).Offset(offset).Find(&categories).Error; err != nil {
		http.Error(w, "Failed to retrieve categories", http.StatusInternalServerError)
		return
	}

	// Encode the categories to JSON and send the response
	json.NewEncoder(w).Encode(categories)
}

// UpdateCategory updates a category by ID
func (s *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category domains.Category
	if err := s.DB.First(&category, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	categoryRequest := CategoryRequest{}

	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	category.Name = categoryRequest.Name
	category.Description = categoryRequest.Description

	if err := s.DB.Save(&category).Error; err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category)
}

// DeleteCategory deletes a category by ID
func (s *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	if err := s.DB.Delete(&domains.Category{}, id).Error; err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Category deleted")
}
