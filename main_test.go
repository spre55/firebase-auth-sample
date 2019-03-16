package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
}

func TestMyHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/my", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(myHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
}

func TestGetEnvVars(t *testing.T) {
	envVars := getEnvVars()
	expected := map[string]string{
		"apiKey":            os.Getenv("API_KEY"),
		"authDomain":        os.Getenv("AUTH_DOMAIN"),
		"databaseURL":       os.Getenv("DATABASE_URL"),
		"projectId":         os.Getenv("PROJECT_ID"),
		"storageBucket":     os.Getenv("STORAGE_BUCKET"),
		"messagingSenderId": os.Getenv("MESSAGING_SENDER_ID"),
	}

	for key, value := range expected {
		if envVars[key] != value {
			t.Errorf(
				"unexpected env: got (%v) want (%v)",
				envVars[key],
				value,
			)
		}
	}
}

func TestIndexHandlerNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/404", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
}

func TestMyHandlerNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/404", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
}
