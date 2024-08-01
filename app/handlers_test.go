package app

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Create a default logger for testing
func createLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil)) // Replace with your actual logger setup
}

func TestGetHandler(t *testing.T) {
	logger := createLogger()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler := getHandler(logger)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("getHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "get request"
	body := rr.Body.String()
	if body != expected {
		t.Errorf("getHandler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestPostHandler(t *testing.T) {
	logger := createLogger()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()

	handler := postHandler(logger)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("postHandler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	expected := "post request"
	body := rr.Body.String()
	if body != expected {
		t.Errorf("postHandler returned unexpected body: got %v want %v", body, expected)
	}
}
