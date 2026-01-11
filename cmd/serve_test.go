package cmd

import (
	"bytes"
	"encoding/json"
	"go-task-master/internal/db"
	"go-task-master/internal/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTasksHandler(t *testing.T) {
	// Initialize the database for testing
	if err := db.InitDB(":memory:"); err != nil {
		t.Fatalf("Failed to initialize in-memory database: %v", err)
	}
	defer db.DB.Close()

	// Seed the database for GET test
	mockTasks := []models.Task{
		{ID: 1, Title: "Test Task 1", Description: "Description 1", Completed: false},
		{ID: 2, Title: "Test Task 2", Description: "Description 2", Completed: true},
	}

	for _, task := range mockTasks {
		// Note: The ID is auto-incremented by the DB, so we don't insert it.
		_, err := db.DB.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
		if err != nil {
			t.Fatalf("Failed to insert mock data: %v", err)
		}
	}

	t.Run("GET tasks", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/tasks/", nil)
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

		var actualTasks []models.Task
		if err := json.NewDecoder(rr.Body).Decode(&actualTasks); err != nil {
			t.Fatalf("could not decode response: %v", err)
		}

		if !reflect.DeepEqual(actualTasks, mockTasks) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				actualTasks, mockTasks)
		}
	})

	t.Run("POST task", func(t *testing.T) {
		newTask := models.Task{
			Title:       "New Task",
			Description: "New Description",
			Completed:   false,
		}
		body, _ := json.Marshal(newTask)
		req, err := http.NewRequest("POST", "/tasks/", bytes.NewBuffer(body))
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

		var createdTask models.Task
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

	t.Run("PUT task", func(t *testing.T) {
		// Using existing task with ID 1 from the mock data
		updatedTask := models.Task{
			ID:          1,
			Title:       "Updated Task 1",
			Description: "Updated Description 1",
			Completed:   true,
		}
		body, _ := json.Marshal(updatedTask)
		req, err := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
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

		// Verify the task was updated in the database
		var task models.Task
		err = db.DB.QueryRow("SELECT id, title, description, completed FROM tasks WHERE id = 1").Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			t.Fatalf("Failed to retrieve updated task from DB: %v", err)
		}

		if !reflect.DeepEqual(task, updatedTask) {
			t.Errorf("task was not updated correctly in DB: got %+v want %+v", task, updatedTask)
		}
	})

	t.Run("DELETE task", func(t *testing.T) {
		// Let's delete the task with ID 2
		req, err := http.NewRequest("DELETE", "/tasks/2", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tasksHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}

		// Verify the task was deleted from the database
		var count int
		err = db.DB.QueryRow("SELECT COUNT(*) FROM tasks WHERE id = 2").Scan(&count)
		if err != nil {
			t.Fatalf("Failed to query DB for deleted task: %v", err)
		}
		if count > 0 {
			t.Errorf("task was not deleted from DB")
		}
	})
}
