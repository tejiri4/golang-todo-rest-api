package routes

import (
	"github.com/gorilla/mux"
	"github.com/tejiri4/golang-todo-rest-api/controller"
)

func Routes (router *mux.Router) {
	router.HandleFunc("/todos", todo.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", todo.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos/{id}", todo.PatchTodo).Methods("PATCH")
	router.HandleFunc("/todos/{id}", todo.GetTodo).Methods("GET")
	router.HandleFunc("/todos", todo.CreateTodo).Methods("POST")
}