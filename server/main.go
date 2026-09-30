package main

import (
	"net/http"
	"encoding/json"
	"time"

	"github.com/so1ove1/dev_tech_2/daysfornewyear"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/days-for-new-year", handler)
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date");

	var d time.Time
	if date == "" {
		d = time.Now()
	} else {
		parsed, err := time.Parse("2006-01-02", date)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Некорректный формат даты, ожидается ГГГГ-ММ-ДД"})
			return
		}
		d = parsed
	}

	result := daysfornewyear.DaysUntilNewYear(d)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"days-for-new-year": result})
}