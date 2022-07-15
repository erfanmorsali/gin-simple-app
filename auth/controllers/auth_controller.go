package controllers

import (
	"github.com/erfanmorsali/gin-simple-app.git/auth/dtos"
	"github.com/erfanmorsali/gin-simple-app.git/auth/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService interfaces.AuthService
	jwtService  interfaces.JwtService
}

func NewAuthController(authService interfaces.AuthService, jwtService interfaces.JwtService) *AuthController {
	return &AuthController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c AuthController) Login(context *gin.Context) {
	var loginInput dtos.LoginIn

	err := context.ShouldBindJSON(&loginInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isAuthenticated := c.authService.Login(loginInput.Email, loginInput.Password)
	if !isAuthenticated {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User Credentials",
		})
		return
	}

	token := c.jwtService.GenerateToken(loginInput.Email, true)
	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return
}

func (c AuthController) Register(context *gin.Context) {
	var registerInput dtos.RegisterIn
	err := context.ShouldBindJSON(&registerInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	registerOut, err := c.authService.Register(registerInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, registerOut)
	return
}
