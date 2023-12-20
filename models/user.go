package models

const UserId = "user-id"

type RegisterParams struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm"`
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
}
