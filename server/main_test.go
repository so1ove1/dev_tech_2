package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_ValidDate(t *testing.T) {
	req := httptest.NewRequest("GET", "/days-for-new-year?date=2026-03-01", nil)
	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Ожидается статус %d, получен %d", http.StatusOK, rec.Code)
	}

	var response map[string]int
	err := json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		t.Errorf("Ошибка расшифровки ответа: %v", err)
	}

	if response["days-for-new-year"] != 306 {
		t.Errorf("Ожидается 306 дней до нового года, получено %d", response["days-for-new-year"])
	}
}

func TestHandler_InvalidDate(t *testing.T) {
	req := httptest.NewRequest("GET", "/days-for-new-year?date=2026-90-52", nil)
	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Ожидается статус %d, получен %d", http.StatusBadRequest, rec.Code)
	}
}

func TestHandler_NoDate(t *testing.T) {
	req := httptest.NewRequest("GET", "/days-for-new-year", nil)
	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Ожидается статус %d, получен %d", http.StatusOK, rec.Code)
	}

	var response map[string]int
	err := json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		t.Errorf("Ошибка расшифровки ответа: %v", err)
	}
}
