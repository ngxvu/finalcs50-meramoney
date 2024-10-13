package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"meramoney/backend/infrastructure/domains"
	"net/http"
	"strconv"
)

type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateCategory creates a new category
func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value("id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	var categoryRequest CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if the category already exists for the user
	var existingCategory domains.Category
	if err := s.DB.Where("name = ? AND user_id = ?", categoryRequest.Name, userID).First(&existingCategory).Error; err == nil {
		http.Error(w, "Category already exists", http.StatusConflict)
		return
	}

	var category domains.Category
	category.Name = categoryRequest.Name
	category.Description = categoryRequest.Description
	category.UserID = userID

	if err := s.DB.Create(&category).Error; err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

// GetCategory retrieves a category by ID
func (s *Server) GetCategory(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value("id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category domains.Category
	if err := s.DB.Where("user_id = ?", userID).First(&category, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(category)
}

// GetAllCategories retrieves all categories
func (s *Server) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value("id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

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
	if err := s.DB.Limit(pageSize).Offset(offset).Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		http.Error(w, "Failed to retrieve categories", http.StatusInternalServerError)
		return
	}

	// Encode the categories to JSON and send the response
	json.NewEncoder(w).Encode(categories)
}

// UpdateCategory updates a category by ID
func (s *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value("id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category domains.Category
	if err := s.DB.Where("user_id = ?", userID).First(&category, id).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	categoryRequest := CategoryRequest{}

	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check for duplicate category name
	var existingCategory domains.Category
	if err := s.DB.Where("name = ? AND user_id = ? AND id != ?", categoryRequest.Name, userID, id).First(&existingCategory).Error; err == nil {
		http.Error(w, "Category name already exists", http.StatusConflict)
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

	userID, ok := r.Context().Value("id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// Check if the category is used in transactions
	var transactionCount int64
	if err := s.DB.Model(&domains.Transaction{}).Where("category_id = ? AND user_id = ?", id, userID).Count(&transactionCount).Error; err != nil {
		http.Error(w, "Failed to check category usage", http.StatusInternalServerError)
		return
	}

	if transactionCount > 0 {
		http.Error(w, "Category is in use and cannot be deleted", http.StatusConflict)
		return
	}

	if err := s.DB.Delete(&domains.Category{}, id).Error; err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Category deleted")
}
