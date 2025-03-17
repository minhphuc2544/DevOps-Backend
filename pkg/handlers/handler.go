package handlers

import (
    "net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("GET request handled"))
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("POST request handled"))
}