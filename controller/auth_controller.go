package controller

import (
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Registered"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"token": "mock-jwt-token"})
}
