package server

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Err    string
	Status int
}

func (e ApiError) Error() string {
	return e.Err
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

// Centralized error handling
// This function takes an ApiFunc as input and returns an http.HandlerFunc.
// It acts as a middleware that handles errors returned by the ApiFunc.
// If the ApiFunc returns an ApiError, it writes the error as JSON response.
// Otherwise, it writes a generic internal server error response.
func MakeHTTPHandler(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(ApiError); ok {
				WriteJSON(w, e.Status, e)
				return
			}
			WriteJSON(w, http.StatusInternalServerError, ApiError{Err: "internal server error", Status: http.StatusInternalServerError})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
