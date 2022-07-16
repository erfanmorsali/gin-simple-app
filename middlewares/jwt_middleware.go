package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/erfanmorsali/gin-simple-app.git/auth/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthenticationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := services.NewJwtService().ValidateToken(tokenString)

		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			context.Set("userCredentials", claims)
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token",
			})
		}
	}
}
