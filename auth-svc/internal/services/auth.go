package services

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/dto"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/entities"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/repository"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/shared"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CreateUser(ctx context.Context, user dto.AuthRequest) (*dto.AuthResponse, error)
	Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error)
	Claim(ctx context.Context, tokenString string) (*dto.TokenResponse, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (a *authService) CreateUser(ctx context.Context, user dto.AuthRequest) (*dto.AuthResponse, error) {
	password := utils.GenerateRandomString(4)

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = a.repo.CreateUser(ctx, entities.User{
		Model: entities.Model{
			ID: uuid.New().String(),
		},
		Name:     user.Name,
		Phone:    user.Phone,
		Role:     user.Role,
		Password: string(hashedPwd),
	})
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Name:     user.Name,
		Phone:    user.Phone,
		Role:     user.Role,
		Password: password,
	}, nil
}

func (a *authService) Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := a.repo.FetchUserByPhone(ctx, request.Phone)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}

	// Create JWT token
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &dto.TokenResponse{
		Name:  user.Name,
		Phone: user.Phone,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(shared.JWTKey))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: tokenString,
	}, nil
}

func (a *authService) Claim(ctx context.Context, tokenString string) (*dto.TokenResponse, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.TokenResponse{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(shared.JWTKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*dto.TokenResponse); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
