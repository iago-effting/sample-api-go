package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"iago-effting/api-example/pkg/authentication"
	authDb "iago-effting/api-example/pkg/storage/database/authentication"
)

func AuthHandler(ctx *gin.Context) {
	var credentials authentication.Credentials
	var authRepo authentication.Repository = authDb.Repo()
	var jwtService = authentication.JWTAuthService()

	if err := ctx.BindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isValid, err := authRepo.CheckCredentials(ctx, credentials)
	if err != nil || !isValid {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Credentials are not valid",
		})
		return
	}

	token := jwtService.GenerateToken(credentials.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
