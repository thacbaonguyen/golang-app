package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-ginapp/models"
	"time"
)

type JWTServiceImpl struct {
	secretKey []byte
	expires   int64
}

func NewJWTServiceImpl(secretKey string, expires int64) JWTService {
	return &JWTServiceImpl{
		secretKey: []byte(secretKey),
		expires:   expires,
	}
}

func (J *JWTServiceImpl) GenerateToken(user models.User) (string, error) {
	claims := JWTClaims{
		UserId: user.ID,
		RoleId: user.RoleId,
		Role:   user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second + time.Duration(J.expires))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(J.secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (J *JWTServiceImpl) ValidateToken(tokenString string) (*JWTClaims, error) {
	//TODO implement me
	token, err := jwt.ParseWithClaims(tokenString,
		&JWTClaims{},
		// get secret
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // ep kieu token.Method ve HMAC -- _ bo qua gia tri ep, chi quan tam den ok
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return J.secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims) // ep kieu token.Claims sang JWTClaims
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
