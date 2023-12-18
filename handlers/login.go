package handlers

import (
	"encoding/json"
	"goHttpServer/models"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 解析请求参数
	var body models.LoginParams
	rawBody := json.NewDecoder(r.Body)
	if err := rawBody.Decode(&body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	}
	if body.Email == "" || body.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("email && password is required!"))
		if err != nil {
			panic(err)
		}
		return
	}
	if body.Email == "obama@gmail.com" && body.Password == "obama@gmail.com" {
		cookie := &http.Cookie{Name: "haha", Value: "obama", Path: "/"}
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("login successfully"))
		if err != nil {
			panic(err)
		}
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("your email or password is invalid! please try again."))
	if err != nil {
		panic(err)
	}
	return
}
