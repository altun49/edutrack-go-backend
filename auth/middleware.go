package auth

import "github.com/gin-gonic/gin"

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// JWT validation logic
	}
}

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Role validation logic
	}
}
