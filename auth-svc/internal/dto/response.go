package dto

import "github.com/dgrijalva/jwt-go"

type AuthResponse struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type TokenResponse struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type ErrorReply struct {
	Error ErrorMessage `json:"error"`
}

func ReplyError(message string) ErrorReply {
	return ErrorReply{
		Error: ErrorMessage{
			Message: message,
		},
	}
}
