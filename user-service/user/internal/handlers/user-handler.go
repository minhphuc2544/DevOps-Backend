package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
}


func (h* Handler) CreateNewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the request body to get user details
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here you would typically save the user to a database and hash the password
	user.CreatedAt = time.Now().Format(time.RFC3339)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	if err := h.db.Create(&user).Error; err != nil {
		http.Error(w, "Failed to write data into database", http.StatusInternalServerError)
		return
	}

	id := user.ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Sucessfully created user",
		"id":      id,
	})
}