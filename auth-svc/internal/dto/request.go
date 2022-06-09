package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type AuthRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

func (a *AuthRequest) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Phone, validation.Required),
		validation.Field(&a.Role, validation.Required),
	)
}

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (l *LoginRequest) Validate() error {
	return validation.ValidateStruct(l,
		validation.Field(&l.Phone, validation.Required),
		validation.Field(&l.Password, validation.Required),
	)
}

type TokenRequest struct {
	Token string `json:"token"`
}

func (t *TokenRequest) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Token, validation.Required),
	)
}
