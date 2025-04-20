package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
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
		http.Error(w, "Lỗi hash mật khẩu", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	stmt, err := h.db.Prepare("INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Lỗi chuẩn bị câu lệnh", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Username, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		http.Error(w, "Lỗi khi ghi dữ liệu", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Tạo người dùng thành công",
		"id":      id,
	})
}