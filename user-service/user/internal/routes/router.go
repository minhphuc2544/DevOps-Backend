package routes

import (
	"database/sql"

	"github.com/julienschmidt/httprouter"
	"github.com/minhphuc2544/DevOps-Backend/user-service/user/internal/handlers"
)

func SetupRoutes(db *sql.DB) *httprouter.Router {
	router := httprouter.New()
	h := handlers.NewHandler(db)
	// Define your routes here
	router.POST("/users", h.CreateNewUser)

	return router
}