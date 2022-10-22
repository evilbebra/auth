package controller

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, you are successfully authenticated!"))
	return
}
