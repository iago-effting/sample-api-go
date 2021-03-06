package midlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"iago-effting/api-example/pkg/authentication"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		jwtService := authentication.JWTAuthService()

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("identity", claims["identity"])
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}
