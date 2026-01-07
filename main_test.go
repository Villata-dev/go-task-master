package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTasksHandler(t *testing.T) {
	// Initialize the database for testing
	if err := initDB(":memory:"); err != nil {
		t.Fatalf("Failed to initialize in-memory database: %v", err)
	}
	defer db.Close()

	// Seed the database for GET test
	mockTasks := []Task{
		{ID: 1, Title: "Test Task 1", Description: "Description 1", Completed: false},
		{ID: 2, Title: "Test Task 2", Description: "Description 2", Completed: true},
	}

	for _, task := range mockTasks {
		// Note: The ID is auto-incremented by the DB, so we don't insert it.
		_, err := db.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
		if err != nil {
			t.Fatalf("Failed to insert mock data: %v", err)
		}
	}

	t.Run("GET tasks", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/tasks", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tasksHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		var actualTasks []Task
		if err := json.NewDecoder(rr.Body).Decode(&actualTasks); err != nil {
			t.Fatalf("could not decode response: %v", err)
		}

		if !reflect.DeepEqual(actualTasks, mockTasks) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				actualTasks, mockTasks)
		}
	})

	t.Run("POST task", func(t *testing.T) {
		newTask := Task{
			Title:       "New Task",
			Description: "New Description",
			Completed:   false,
		}
		body, _ := json.Marshal(newTask)
		req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tasksHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}

		var createdTask Task
		if err := json.NewDecoder(rr.Body).Decode(&createdTask); err != nil {
			t.Fatalf("could not decode response: %v", err)
		}

		if createdTask.ID == 0 {
			t.Errorf("expected created task to have an ID, but it was 0")
		}

		// Check if the title and other fields match
		if createdTask.Title != newTask.Title || createdTask.Description != newTask.Description || createdTask.Completed != newTask.Completed {
			t.Errorf("handler returned unexpected body for created task: got %+v want %+v",
				createdTask, newTask)
		}
	})
}
