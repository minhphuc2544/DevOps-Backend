package main

import (
    "net/http"
    "log"
	"github.com/minhphuc2544/DevOps-Backend/user-service/auth/internal/handlers"
    "github.com/julienschmidt/httprouter"
)



func main() {
    router := httprouter.New()
    router.GET("/", handlers.Index)
	router.GET("/hello/:name", handlers.Hello)

	log.Println("Starting server on :8080...")
	// Start the server on port 8080
    log.Fatal(http.ListenAndServe(":8080", router))
}