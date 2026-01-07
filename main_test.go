package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetTasks(t *testing.T) {
	initTasks()

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTasks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedTasks := tasks
	var actualTasks []Task
	if err := json.NewDecoder(rr.Body).Decode(&actualTasks); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if !reflect.DeepEqual(actualTasks, expectedTasks) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actualTasks, expectedTasks)
	}
}
