package user

import (
	"go-ginapp/data/request"
	"go-ginapp/data/response"
)

type UserService interface {
	GetAllUser() ([]response.UserResponse, error)
	GetUserByID(id uint) (response.UserResponse, error)
	UpdateUser(id uint, userUpdateReq request.UserUpdateReq) (response.UserResponse, error)
	UpdatePassword(id uint, passwordRequest request.UpdatePasswordRequest) error
	DeleteUser(id uint) error
}
