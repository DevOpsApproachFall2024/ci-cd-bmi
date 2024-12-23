package handlers

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCalculateBMI(t *testing.T) {
    reqBody := `{"height":1.75,"weight":68}`
    req := httptest.NewRequest(http.MethodPost, "/calculate-bmi", bytes.NewBuffer([]byte(reqBody)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    CalculateBMI(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("Expected status OK but got %v", w.Code)
    }
}
