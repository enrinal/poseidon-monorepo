package repository

import (
	"context"

	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/entities"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user entities.User) error
	FetchUserByPhone(ctx context.Context, phone string) (*entities.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

type authRepository struct {
	db *gorm.DB
}

func (a *authRepository) CreateUser(ctx context.Context, user entities.User) error {
	return a.db.WithContext(ctx).Create(&user).Error
}

func (a *authRepository) FetchUserByPhone(ctx context.Context, phone string) (*entities.User, error) {
	var user entities.User
	err := a.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	return &user, err
}
