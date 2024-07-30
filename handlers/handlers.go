package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/assaidy/task-manager-api/repo"
	"github.com/gorilla/mux"
)

// writeErrorResponse writes an error message to the HTTP response
func writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}

// encodeResponse encodes the response to JSON and writes it to the HTTP response
func encodeResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
		writeErrorResponse(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// HandleGetAllTasks handles GET requests to retrieve all tasks
func HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	allTasks := repo.GetAllTasks()
	encodeResponse(w, allTasks, http.StatusOK)
}

// HandleCreateNewTask handles POST requests to create a new task
func HandleCreateNewTask(w http.ResponseWriter, r *http.Request) {
	contentQuery := r.URL.Query().Get("content")
	if contentQuery == "" {
		writeErrorResponse(w, "Content is required", http.StatusBadRequest)
		return
	}

	t := repo.CreateTask(contentQuery)
	encodeResponse(w, t, http.StatusCreated)
}

// HandleGetTask handles GET requests to retrieve a task by ID
func HandleGetTask(w http.ResponseWriter, r *http.Request) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		writeErrorResponse(w, "Task ID is missing in parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeErrorResponse(w, "Id is not a valid number", http.StatusBadRequest)
		return
	}

	t, err := repo.GetTask(id)
	if err != nil {
		writeErrorResponse(w, "No task with such id", http.StatusNotFound)
		return
	}

	encodeResponse(w, t, http.StatusOK)
}

// HandleToggleComplete handles POST requests to toggle the completion status of a task
func HandleToggleComplete(w http.ResponseWriter, r *http.Request) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		writeErrorResponse(w, "Task ID is missing in parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeErrorResponse(w, "Id is not a valid number", http.StatusBadRequest)
		return
	}

	t, err := repo.ToggleTaskComplete(id)
	if err != nil {
		writeErrorResponse(w, "No task with such id", http.StatusNotFound)
		return
	}

	encodeResponse(w, t, http.StatusOK)
}

// HandleDeleteTask handles DELETE requests to delete a task by ID
func HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		writeErrorResponse(w, "Task ID is missing in parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeErrorResponse(w, "Id is not a valid number", http.StatusBadRequest)
		return
	}

	err = repo.DeleteTask(id)
	if err != nil {
		writeErrorResponse(w, "No task with such id", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// HandleUpdateTask handles PUT requests to update a task by ID
func HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		writeErrorResponse(w, "Task ID is missing in parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeErrorResponse(w, "Id is not a valid number", http.StatusBadRequest)
		return
	}

	contentQuery := r.URL.Query().Get("content")
	if contentQuery == "" {
		writeErrorResponse(w, "Content is required", http.StatusBadRequest)
		return
	}

	t, err := repo.UpdateTask(id, contentQuery)
	if err != nil {
		writeErrorResponse(w, "No task with such id", http.StatusNotFound)
		return
	}

	encodeResponse(w, t, http.StatusOK)
}

// HandleDeleteAllTasks handles DELETE requests to delete all tasks
func HandleDeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	repo.DeleteAllTasks()
	w.WriteHeader(http.StatusNoContent)
}
