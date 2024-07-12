package request

type RegisterRequest struct {
	Mobile               string `json:"mobile"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	Code                 string `json:"code"`
}
