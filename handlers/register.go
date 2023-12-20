package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"userSystem/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if userId := r.Header.Get(models.UserId); userId != "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(models.HaveAccountError))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	// 解析请求参数
	var requestBody = models.RegisterParams{}
	rawBody := json.NewDecoder(r.Body)
	if err := rawBody.Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 查询注册的email是否在数据库中存在
	/*

	 */

	// 检查密码一致性
	if requestBody.Password != requestBody.ConfirmPassword {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(models.PasswordNotConfirmError))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	// 弱密码检查：password必须以大小写字母开头，长度不小于8位，并且至少包含一个符号
	if ok, _ := regexp.MatchString("^(?=.*[!@#$%^&*()-=_+{};':\",.<>/?])(?=.*\\d)[a-zA-Z\\d].{6,}$", requestBody.Password); !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(models.PasswordDoesMeetRuleError))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	// 检查完成，写入db
	/*

	 */
}
