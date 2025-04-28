package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"go-ginapp/models"
)

type JWTClaims struct {
	UserId uint   `json:"userId"`
	RoleId uint   `json:"roleId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(user models.User) (string, error)
	ValidateToken(token string) (*JWTClaims, error)
}
