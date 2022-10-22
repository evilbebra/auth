package main

import (
	"github.com/evilbebra/auth/cmd/auth-app/internal/controller"
	"github.com/evilbebra/auth/config"
	"github.com/evilbebra/auth/internal/auth"
	"log"
	"net/http"
)

func main() {
	conf := config.Init("./config/config.yaml")

	authService := auth.NewAuthService(conf.Auth.SigningKey, conf.Auth.TokenTTL)
	authHandler := controller.NewAuthHandler(authService)

	mux := http.NewServeMux()
	mux.HandleFunc("/", authHandler.HandleAuth)

	log.Fatal(http.ListenAndServe(":"+conf.AuthApp.Port, mux))
}
