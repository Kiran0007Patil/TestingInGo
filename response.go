package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func failureResponse(w http.ResponseWriter, message string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(struct {
		Error   string
		IsError bool `json:"error"`
	}{
		Error:   message,
		IsError: true,
	})
}

func successResponse(w http.ResponseWriter, j []byte) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if bytes.Equal(j, []byte("null")) {
		fmt.Fprintf(w, "%s", "{}")
	} else {
		fmt.Fprintf(w, "%s", j)
	}
}
