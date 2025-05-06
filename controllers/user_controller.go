package controllers

import (
	"github.com/gin-gonic/gin"
	"go-ginapp/data/request"
	"go-ginapp/services/user"
	"go-ginapp/utils"
	"net/http"
	"strconv"
)

type UserController interface {
	GetAllUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUser(c *gin.Context)
	ChangePassword(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetCurrentUser(c *gin.Context)
}

type userController struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) UserController {
	return &userController{userService}
}

func (u *userController) GetAllUsers(c *gin.Context) {
	response, err := u.userService.GetAllUser()
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Get all user failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Get all user success", response)
}

func (u *userController) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user id", err)
		return
	}
	response, err := u.userService.GetUserByID(uint(id))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Get all user failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Get all user success", response)
}

func (u *userController) UpdateUser(c *gin.Context) {
	id := c.GetUint("userId")

	var userUpdate request.UserUpdateReq
	if err := utils.ValidateStruct(&userUpdate); err != nil {
		errors := utils.GetValidationErrors(err)
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", errors)
		return
	}
	response, err := u.userService.UpdateUser(id, userUpdate)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Update failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Update user success", response)
}

func (u *userController) ChangePassword(c *gin.Context) {
	id := c.GetUint("userId")
	var changePassRq request.UpdatePasswordRequest
	if err := utils.ValidateStruct(&changePassRq); err != nil {
		errors := utils.GetValidationErrors(err)
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", errors)
		return
	}
	if err := u.userService.UpdatePassword(uint(id), changePassRq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Update password failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Update password success", nil)
}

func (u *userController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user id", err)
		return
	}
	if err = u.userService.DeleteUser(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Delete user failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Delete user success", nil)

}

func (u *userController) GetCurrentUser(c *gin.Context) {
	id := c.GetUint("userId")
	response, err := u.userService.GetUserByID(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Get your info failed", err)
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Get your profile success", response)
}
