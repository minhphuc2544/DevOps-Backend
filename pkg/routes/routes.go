package routes

import (
    "github.com/gorilla/mux"
    "my-go-backend/pkg/handlers"
)

func SetupRoutes(r *mux.Router) {
    // base on the routes, the handlers will be called
    r.HandleFunc("/api/resource", handlers.HandleGet).Methods("GET") //This will call the HandleGet function from the handlers package
    r.HandleFunc("/api/resource", handlers.HandlePost).Methods("POST") //This will call the HandlePost function from the handlers package
}