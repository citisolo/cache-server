package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

// ParseBody parses body of POST request
func ParseBody( body io.ReadCloser, data interface{} ) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&data)
	if err != nil {
		return err
	}
	defer body.Close()
	return nil
}