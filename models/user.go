package models

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const UserId = "user-id"
