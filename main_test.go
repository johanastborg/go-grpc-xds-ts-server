package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Response{Message: "Hello, Bazel-built Go REST API!!!"}
	var actual Response
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if actual.Message != expected.Message {
		t.Errorf("handler returned unexpected message: got %v want %v",
			actual.Message, expected.Message)
	}
}

func TestHelloHandlerMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}
