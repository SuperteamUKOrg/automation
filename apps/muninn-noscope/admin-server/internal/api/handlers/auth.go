package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"admin-server/internal/database"
)

type AuthHandler struct {
	queries *database.Queries
	logger  *log.Logger
}

func NewAuthHandler(q *database.Queries, l *log.Logger) *AuthHandler {
	return &AuthHandler{
		queries: q,
		logger:  l,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	reqUserID := r.Header.Get("x-user-secret")
	secret := os.Getenv("ADMIN_SECRET")
	if reqUserID == "" || reqUserID != secret {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	muninn_jwt := os.Getenv("MUNINN_JWT")
	control_secret := os.Getenv("CONTROL_SECRET_KEY")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"muninn_jwt": muninn_jwt,
		"control_secret": control_secret,
	})
}