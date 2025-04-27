package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/minhphuc2544/DevOps-Backend/user-service/user/internal/utils"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}


func (h* Handler) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get all users from the database
	var users []User
	if err := h.db.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the list of users as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users": users,
		"count": len(users),
		"message": "Get all users successfully",
	})
}

func (h *Handler) GetInfoByUsername(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get the username from the URL parameters
	claims, err := utils.VerifyJWT(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}
	fmt.Println("Logged in user ID:", claims["user_id"])

	var requestBody struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, "Invalid JSON body", http.StatusBadRequest)
    }
	username := requestBody.Username
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Get the user from the database
	var user User
	if err := h.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the user as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"fullname": user.Fullname,
		"email":    user.Email,
		"created_at": user.CreatedAt,
	})
}