package controller

import (
	"github.com/evilbebra/auth/internal/auth"
	"github.com/evilbebra/auth/internal/types"
	"log"
	"net/http"
)

type AuthHandler struct {
	auth auth.Auth
}

func NewAuthHandler(auth auth.Auth) *AuthHandler {
	return &AuthHandler{
		auth: auth,
	}
}

func (h *AuthHandler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	user := &types.User{
		ID:       1,
		Email:    "evilbebra@gmail.com",
		Password: "password",
	}
	token, err := h.auth.GenerateJWTToken(user)
	if err != nil {
		log.Printf("HandleAuth - ERROR %v\n", err)
	}
	w.Write([]byte(token))
}
