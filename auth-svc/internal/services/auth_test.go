package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/dto"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/entities"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/shared"
	mock "github.com/enrinal/poseidon-monorepo/auth-svc/mock/repository"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthService_CreateUser(t *testing.T) {
	Convey("Create User", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoAuth := mock.NewMockAuthRepository(ctrl)

		authService := NewAuthService(repoAuth)

		dtoAuth := dto.AuthRequest{
			Name:  "John Doe",
			Phone: "62812341234",
			Role:  "ADMIN",
		}

		Convey("Positive Scenario", func() {
			Convey("Should return success with user", func() {
				repoAuth.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
				res, err := authService.CreateUser(context.Background(), dtoAuth)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})

		Convey("Negative Scenario", func() {
			Convey("Should return error", func() {
				repoAuth.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(errors.New("error"))
				res, err := authService.CreateUser(context.Background(), dtoAuth)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
	})
}

func TestAuthService_Login(t *testing.T) {
	Convey("Login", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoAuth := mock.NewMockAuthRepository(ctrl)

		authService := NewAuthService(repoAuth)

		dtoLogin := dto.LoginRequest{
			Phone:    "62812341234",
			Password: "123456",
		}

		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(dtoLogin.Password), bcrypt.DefaultCost)

		auth := entities.User{
			Model: entities.Model{
				ID: "uuidv4",
			},
			Name:     "John Doe",
			Phone:    "62812341234",
			Role:     "ADMIN",
			Password: string(hashedPwd),
		}

		Convey("Positive Scenario", func() {
			Convey("Should return success with token", func() {
				repoAuth.EXPECT().FetchUserByPhone(gomock.Any(), dtoLogin.Phone).Return(&auth, nil)
				res, err := authService.Login(context.Background(), dtoLogin)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("Should return error", func() {
				repoAuth.EXPECT().FetchUserByPhone(gomock.Any(), dtoLogin.Phone).Return(nil, errors.New("error"))
				res, err := authService.Login(context.Background(), dtoLogin)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
	})
}

func TestAuthService_Claim(t *testing.T) {
	Convey("Claim", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoAuth := mock.NewMockAuthRepository(ctrl)

		authService := NewAuthService(repoAuth)

		Convey("Positive Scenario", func() {
			Convey("Should return success with right token", func() {
				expirationTime := time.Now().Add(60 * time.Minute)
				claims := &dto.TokenResponse{
					Name:  "John Doe",
					Phone: "62812341234",
					Role:  "ADMIN",
					StandardClaims: jwt.StandardClaims{
						ExpiresAt: expirationTime.Unix(),
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, _ := token.SignedString([]byte(shared.JWTKey))

				res, err := authService.Claim(context.Background(), tokenString)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
				So(res.Name, ShouldEqual, "John Doe")
				So(res.Phone, ShouldEqual, "62812341234")
				So(res.Role, ShouldEqual, "ADMIN")
			})
		})
		Convey("Negative Scenario", func() {
			Convey("Should return error", func() {
				tokenString := "invalid token"
				res, err := authService.Claim(context.Background(), tokenString)
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
	})
}
