package main

import (
	"net/http"
	"net/http/httptest"
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

	expected := "Sign Up!"
	if rr.Body.String() != expected {
		t.Errorf(
			"unexpected body: got (%v) want (%v)",
			rr.Body.String(),
			expected,
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

	expected := "My Page!"
	if rr.Body.String() != expected {
		t.Errorf(
			"unexpected body: got (%v) want (%v)",
			rr.Body.String(),
			expected,
		)
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