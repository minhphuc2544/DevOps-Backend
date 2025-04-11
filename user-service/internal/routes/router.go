package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/minhphuc2544/DevOps-Backend/user-service/internal/handlers"
)

func SetupRoutes() *httprouter.Router {
	router := httprouter.New()

	// Define your routes here
	router.GET("/", handlers.Index)
	router.GET("/hello/:name", handlers.Hello)

	return router
}