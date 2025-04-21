package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
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


	// Validate the user details
	if user.Username == "" || user.Fullname == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	// Check if email is valid
	if !strings.Contains(user.Email, "@") {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}	
	// Check if password is strong enough
	if len(user.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now().Format(time.RFC3339)

	// write the user to the database
	if err := h.db.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "username"){
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
		if strings.Contains(err.Error(), "email") {
			http.Error(w, "Email already exists", http.StatusConflict)
			return	
		}
	}

	id := user.ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Sucessfully created user",
		"id":      id,
	})
}