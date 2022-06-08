package dto

type AuthRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type TokenRequest struct {
	Token string `json:"token"`
}
