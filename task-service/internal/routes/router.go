package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/minhphuc2544/DevOps-Backend/task-service/internal/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *httprouter.Router {
	router := httprouter.New()
	h := handlers.NewHandler(db)
	// Define your routes here
	router.GET("/getAllMusic", h.GetAllMusic) // Get all music
	router.POST("/uploadMusic", h.UploadMusic) // Upload music
	router.POST("/createPlaylist", h.CreatePlaylist) // Create playlist
	router.POST("/addMusicToPlaylist", h.AddMusicToPlaylist) // Add music to playlist
	router.GET("/getUserPlaylist", h.GetUserPlayListByUserId) // Get playlist
	return router
}
