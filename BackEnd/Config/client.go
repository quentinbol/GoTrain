package main

import (
    "context"
	"encoding/json"
	"log"
	"net/http"
    "GoTrain/BackEnd/router"
    "github.com/gorilla/mux"

    _ "github.com/mattn/go-sqlite3"
)

var client, errInit = router.InitializeDB()

func main() {
    if errInit != nil {
        log.Fatalf("failed to initialize database: %v", errInit)
    }
    defer client.Close()

    r := mux.NewRouter()
	r.HandleFunc("/messages", getMessages).Methods("GET")
	r.HandleFunc("/messages", createMessage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	messages, err := router.QueryAllMessages(ctx, client)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Content string `json:"content"`
		UserID  int    `json:"userId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	newMessage, err := router.CreateMessage(ctx, client, data.Content, data.UserID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newMessage)
}