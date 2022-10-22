package controller

import (
	"github.com/evilbebra/auth/internal/auth"
	"log"
	"net/http"
	"strings"
)

type MiddlewareManager struct {
	auth auth.Auth
}

func NewMiddlewareManager(auth auth.Auth) *MiddlewareManager {
	return &MiddlewareManager{
		auth: auth,
	}
}

func (mw *MiddlewareManager) JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		log.Println("--- JWT Auth")

		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			log.Println(" Authorization header is missing")
			http.Error(w, "401 Unauthorized Error", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			http.Error(w, "401 Unauthorized Error", http.StatusUnauthorized)
			return
		}

		if headerParts[0] != "Bearer" {
			http.Error(w, "401 Unauthorized Error", http.StatusUnauthorized)
			return
		}

		claims, err := mw.auth.ValidateToken(headerParts[1])
		if err != nil {
			log.Printf("ERROR: %v\n", err)
			http.Error(w, "401 Unauthorized Error", http.StatusUnauthorized)
			return
		}

		log.Println("Claims: ", claims)

		next.ServeHTTP(w, req)
	})
}
