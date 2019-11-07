package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tejiri4/golang-todo-rest-api/controller"
)

func homeLink(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/todos", todo.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id:[0-9]+}", todo.GetTodo).Methods("GET")
	router.HandleFunc("/todos", todo.CreateTodo).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}