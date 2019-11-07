package todo

import (
	"net/http"
	"github.com/tejiri4/golang-todo-rest-api/database"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.Todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID, _ := strconv.Atoi(mux.Vars(r)["id"])

	todoFound := false

	for _, todo := range database.Todos {
		if todo.ID == todoID {
			todoFound = true
			json.NewEncoder(w).Encode(todo)
		}
	}

	if !todoFound {
		err := map[string]string{"message": "Todo not found."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reqObj database.Todo
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &reqObj)

	todo := database.Todo{
		ID: len(database.Todos) + 1,
		Todo: reqObj.Todo,
	}

	if errs := todo.Validate(); len(errs) > 0 {
		fmt.Println(todo.Todo)
		err := map[string]interface{}{"message": errs}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		todo.ID = len(database.Todos) + 1
		database.Todos = append(database.Todos, todo)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)
	}
}

