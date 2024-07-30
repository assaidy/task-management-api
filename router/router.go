package router

import (
	"net/http"
    "github.com/assaidy/task-manager-api/handlers"

	"github.com/gorilla/mux"
)

func GetRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/v1/tasks", handlers.HandleGetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/tasks", handlers.HandleCreateNewTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/tasks/{id}", handlers.HandleGetTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/toggleComplete/{id}", handlers.HandleToggleComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/v1/deleteTask/{id}", handlers.HandleDeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/updateTask/{id}", handlers.HandleUpdateTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/v1/deleteAllTasks", handlers.HandleDeleteAllTasks).Methods("DELETE", "OPTIONS")

	return router
}
