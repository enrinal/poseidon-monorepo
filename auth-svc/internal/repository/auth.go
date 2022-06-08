package repository

import (
	"context"

	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/entities"
	"gorm.io/gorm"
)

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user entities.User) error
	FetchUser(ctx context.Context, phone, password string) (*entities.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func (a *authRepository) CreateUser(ctx context.Context, user entities.User) error {
	return a.db.WithContext(ctx).Create(&user).Error
}

func (a *authRepository) FetchUser(ctx context.Context, phone, password string) (*entities.User, error) {
	var user entities.User
	err := a.db.WithContext(ctx).Where("phone = ?", phone).Where("password = ?", password).First(&user).Error
	return &user, err
}
