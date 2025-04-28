package middleware

import (
	"github.com/gin-gonic/gin"
	"go-ginapp/repositories"
	"go-ginapp/services/auth"
	"net/http"
	"strings"
)

type AuthMiddleware interface {
	AuthRequired() gin.HandlerFunc
	AdminRequired() gin.HandlerFunc
}

type authMiddleware struct {
	jwtService auth.JWTService
	userRepo   repositories.UserRepository
}

func NewAuthMiddleware(jwtService auth.JWTService, userRepo repositories.UserRepository) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService, userRepo: userRepo}
}

func (a *authMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
			c.Abort()
			return
		}

		headers := strings.Split(authHeader, "")
		if len(headers) != 2 || headers[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
			c.Abort()
			return
		}

		token := headers[1]
		claims, err := a.jwtService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		_, err = a.userRepo.FindById(claims.UserId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Set("userId", claims.UserId)
		c.Set("roleId", claims.RoleId)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func (a *authMiddleware) AdminRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		// kiem tra auth required truoc, neu bi abort thi return
		a.AuthRequired()(c)
		if c.IsAborted() {
			return
		}

		role := c.GetString("role")
		if role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request admin required"})
			c.Abort()
			return
		}

		c.Next()
	}
}
