package main

import (
	"github.com/evilbebra/auth/cmd/api-app/internal/controller"
	"github.com/evilbebra/auth/config"
	"github.com/evilbebra/auth/internal/auth"
	"log"
	"net/http"
)

func main() {
	conf := config.Init("./config/config.yaml")

	authService := auth.NewAuthService(conf.Auth.SigningKey, conf.Auth.TokenTTL)

	mw := controller.NewMiddlewareManager(authService)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", controller.HelloHandler)
	handler := mw.JWTAuth(mux)

	log.Fatal(http.ListenAndServe(":"+conf.ApiApp.Port, handler))
}
