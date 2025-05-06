package controllers

import (
	"github.com/gin-gonic/gin"
	"go-ginapp/data/request"
	"go-ginapp/services/auth"
	"go-ginapp/utils"
	"net/http"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	authService auth.AuthService
}

func NewAuthController(authService auth.AuthService) AuthController {
	return &authController{authService: authService}
}

func (a *authController) Login(c *gin.Context) {
	var loginRequest request.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if err := utils.ValidateStruct(loginRequest); err != nil {
		errors := utils.GetValidationErrors(err)
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", errors)
		return
	}

	authResponse, err := a.authService.Login(loginRequest)
	if err != nil {
		errors := utils.GetValidationErrors(err)
		utils.ErrorResponse(c, http.StatusUnauthorized, "Authentication failed", errors)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Login successfully", authResponse)
}

func (a *authController) Register(c *gin.Context) {
	var registerRequest request.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if err := utils.ValidateStruct(registerRequest); err != nil {
		errors := utils.GetValidationErrors(err)
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", errors)
		return
	}

	authResponse, err := a.authService.Register(registerRequest)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Register failed", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Register successfully", authResponse)
}
