package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tejiri4/golang-todo-rest-api/database"
	"github.com/gorilla/handlers"
	"github.com/tejiri4/golang-todo-rest-api/routes"
)

func homeLink(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
    database.Db()
	routes.Routes(router)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(router)))
}