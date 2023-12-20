package models

type ErrorCode string

const (
	HaveAccountError          ErrorCode = "You have registered an account, please use your account to log in."
	PasswordNotConfirmError   ErrorCode = "Your password and confirmation password are different. Please confirm and try again."
	PasswordDoesMeetRuleError ErrorCode = "Your password does not meet the rules. The password must start with a letter and contain at least one number and symbol. Please confirm and try again."

	EmailOrPasswordIsRequiredError ErrorCode = "email && password is required!"
	EmailOrPasswordIsInvalidError  ErrorCode = "your email or password is invalid! please try again."
)
