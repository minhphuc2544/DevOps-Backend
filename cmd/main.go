package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "my-go-backend/pkg/routes"
)

func main() {
    router := mux.NewRouter()
    routes.SetupRoutes(router)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}