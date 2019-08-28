package healthchecks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChecker(t *testing.T) {
	checker := NewChecker()
	assert.NotNil(t, checker)
}

func TestSetHealthy(t *testing.T) {
	checker := NewChecker()
	assert.True(t, checker.SetHealthy(true))
	assert.False(t, checker.SetHealthy(false))

}

func TestSetReadiness(t *testing.T) {
	checker := NewChecker()
	assert.True(t, checker.SetReady(true))
	assert.False(t, checker.SetReady(false))

}

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	checker := NewChecker()
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checker.HealthHandlerFunc)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected := "NOT OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHealthHandlerChangeState(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	checker := NewChecker()
	checker.SetHealthy(true)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checker.HealthHandlerFunc)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	checker.SetHealthy(false)
	rr2 := httptest.NewRecorder()
	req2, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr2, req2)
	// Check the status code is what we expect.
	if status2 := rr2.Code; status2 != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status2, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected2 := "NOT OK"
	if rr2.Body.String() != expected2 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr2.Body.String(), expected2)
	}
}

func TestReadinessHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	checker := NewChecker()
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checker.ReadyHandlerFunc)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected := "NOT OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestReadinessHandlerChangeState(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	checker := NewChecker()
	checker.SetReady(true)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checker.ReadyHandlerFunc)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	checker.SetReady(false)
	rr2 := httptest.NewRecorder()
	handler.ServeHTTP(rr2, req)
	// Check the status code is what we expect.
	if status2 := rr2.Code; status2 != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status2, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected2 := "NOT OK"
	if rr2.Body.String() != expected2 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr2.Body.String(), expected2)
	}
}
