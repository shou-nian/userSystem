package handlers

import (
	"encoding/json"
	"net/http"
	"userSystem/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 解析请求参数
	var requestBody models.LoginParams
	rawBody := json.NewDecoder(r.Body)
	if err := rawBody.Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return
	}
	if requestBody.Email == "" || requestBody.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(models.EmailOrPasswordIsRequiredError))
		if err != nil {
			panic(err)
		}
		return
	}
	if requestBody.Email == "obama@gmail.com" && requestBody.Password == "obama@gmail.com" {
		cookie := &http.Cookie{Name: "haha", Value: "obama", Path: "/"}
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("login successfully!"))
		if err != nil {
			panic(err)
		}
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte(models.EmailOrPasswordIsInvalidError))
	if err != nil {
		panic(err)
	}
	return
}
