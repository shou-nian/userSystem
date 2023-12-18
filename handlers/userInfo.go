package handlers

import (
	"net/http"
	"reflect"
	"strconv"
	"userSystem/models"
)

func QueryUserInfo(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get(models.UserId)
	if reflect.DeepEqual(userId, "1") {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(strconv.AppendQuoteToASCII([]byte{}, "email: obama@gmail.com"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		_, err := w.Write(strconv.AppendQuoteToASCII([]byte{}, "user not existing"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
