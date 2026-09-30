package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
		t.Fatalf("Ошибка расшифровки ответа: %v", err)
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

	var response map[string]string
	err := json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Ошибка расшифровки ответа: %v", err)
	}

	if response["error"] != "Некорректный формат даты, ожидается ГГГГ-ММ-ДД" {
		t.Errorf("Ожидается строка \"Некорректный формат даты, ожидается ГГГГ-ММ-ДД\" %v", response["error"])
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
		t.Fatalf("Ошибка расшифровки ответа: %v", err)
	}

	var today = time.Now().Format("2006-01-02")
	req_check := httptest.NewRequest("GET", "/days-for-new-year?date="+today, nil)
	rec_check := httptest.NewRecorder()
	handler(rec_check, req_check)
	var response_check map[string]int
	json.NewDecoder(rec_check.Body).Decode(&response_check)

	if response["days-for-new-year"] != response_check["days-for-new-year"] {
		t.Errorf("Ожидается %d дней до нового года, получено %d", response_check["days-for-new-year"], response["days-for-new-year"])
	}
}
