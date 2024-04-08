package service

type UserLoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}
