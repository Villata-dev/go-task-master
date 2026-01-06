package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Task represents a to-do item.
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks []Task

func initTasks() {
	tasks = []Task{
		{ID: 1, Title: "Complete project proposal", Description: "Write and submit the project proposal for the new client.", Completed: false},
		{ID: 2, Title: "Team meeting", Description: "Attend the weekly team sync-up meeting.", Completed: true},
		{ID: 3, Title: "Code review", Description: "Review the pull request from the junior developer.", Completed: false},
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	initTasks()
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/tasks", getTasks)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
