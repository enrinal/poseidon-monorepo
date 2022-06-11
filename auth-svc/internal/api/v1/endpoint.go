package v1

import (
	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRouterAPIV1(router *gin.RouterGroup, authSvc services.AuthService) {
	userAPI := NewAuthAPI(authSvc)
	router.POST("/users", userAPI.CreateUser)
	router.POST("/users/login", userAPI.Login)
	router.GET("/claims", userAPI.Claims)
}
