package main

import (
	"log"
	"net/http"
	"strconv"
	"userSystem/handlers"
)

func main() {
	http.HandleFunc("/", ServeIndex)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/user-info", handlers.QueryUserInfo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(strconv.AppendQuoteToASCII([]byte{}, "Hello "+r.URL.Query().Get("name")))
	if err != nil {
		return
	}
}
