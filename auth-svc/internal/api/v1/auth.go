package v1

import (
	"net/http"

	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/dto"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/errors"
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthAPI interface {
	Login(c *gin.Context)
	CreateUser(c *gin.Context)
	Claims(c *gin.Context)
}

type authAPI struct {
	authService services.AuthService
}

func NewAuthAPI(authService services.AuthService) AuthAPI {
	return &authAPI{authService: authService}
}

func (a authAPI) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := request.Validate(); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	response, err := a.authService.Login(c, request)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

func (a authAPI) CreateUser(c *gin.Context) {
	var request dto.AuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := request.Validate(); err != nil {
		c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	response, err := a.authService.CreateUser(c, request)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

func (a authAPI) Claims(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.Error(errors.NewError(http.StatusBadRequest, "token is required"))
		return
	}

	response, err := a.authService.Claim(c, token)
	if err != nil {
		c.Error(errors.NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}
