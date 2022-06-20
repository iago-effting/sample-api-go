package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"iago-effting/api-example/pkg/auth"
	"iago-effting/api-example/pkg/storage/database/users"
)

func IndexUser(ctx *gin.Context) {
	var usersRepository auth.Repository = users.Repo()
	var data *[]auth.User

	data, err := usersRepository.All(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func CreateUser(ctx *gin.Context) {
	var createUserRequest CreateUserRequest
	var usersRepository users.Repository = users.Repo()

	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := auth.User{
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	newUser, err := usersRepository.Save(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": newUser,
	})
}

type CreateUserRequest struct {
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeat_password" binding:"required"`
}
