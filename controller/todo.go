package todo

import (
	"net/http"
	"github.com/tejiri4/golang-todo-rest-api/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"context"
	"time"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []database.Todo

	collection := database.Client.Database("todo-golang").Collection("todos")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var todo database.Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	todoID, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])

	var todo database.Todo

	collection := database.Client.Database("todo-golang").Collection("todos")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"_id": todoID}).Decode(&todo)
	if err != nil {
		err := map[string]string{"message": "Todo not found."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reqObj database.Todo
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &reqObj)

	collection := database.Client.Database("todo-golang").Collection("todos")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, reqObj)
	json.NewEncoder(w).Encode(result)
}


func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todoID, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])

	collection := database.Client.Database("todo-golang").Collection("todos")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": todoID})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

    json.NewEncoder(w).Encode(result)
}

func PatchTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqObj database.Todo
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &reqObj)
	
	todoID, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])

	collection := database.Client.Database("todo-golang").Collection("todos")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	updateTodo :=  bson.M{"$set": bson.M{"Todo": reqObj.Todo}}
	updatedResult, err := collection.UpdateOne(ctx, bson.M{"_id": todoID}, updateTodo)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	json.NewEncoder(w).Encode(updatedResult)
	
}