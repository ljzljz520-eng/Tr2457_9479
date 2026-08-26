package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type Health struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}
	json.NewEncoder(w).Encode(Health{Status: "ok", Time: time.Now().UTC()})
}
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
