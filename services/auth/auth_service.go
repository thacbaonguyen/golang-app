package auth

import (
	"go-ginapp/data/request"
	"go-ginapp/data/response"
)

type AuthService interface {
	Login(req request.LoginRequest) (response.AuthResponse, error)
	Register(req request.RegisterRequest) (response.AuthResponse, error)
}
