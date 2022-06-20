package handlers

import (
	"fmt"
	"iago-effting/api-example/pkg/logs"
	"net/http"

	"github.com/gin-gonic/gin"

	"iago-effting/api-example/pkg/account"
	accountDb "iago-effting/api-example/pkg/storage/database/account"
)

func IndexAccount(ctx *gin.Context) {
	var usersRepository account.Repository = accountDb.Repo()
	var data *[]account.User

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

func ViewAccount(ctx *gin.Context) {
	var logger = ctx.MustGet("logger").(logs.Logger)
	var usersRepository account.Repository = accountDb.Repo()
	var data *account.User

	id := ctx.Param("id")
	data, err := usersRepository.Get(ctx, id)

	logger.Info(fmt.Sprintf("Getting view account for ID %s", id))

	if err != nil {
		logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
	}

	logger.Info("Account found")

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func DeleteAccount(ctx *gin.Context) {
	var usersRepository account.Repository = accountDb.Repo()

	id := ctx.Param("id")
	err := usersRepository.Delete(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Deleted successfully",
	})
}

func CreateAccount(ctx *gin.Context) {
	var createUserRequest CreateAccountRequest
	var usersRepository account.Repository = accountDb.Repo()

	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := account.User{
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

type CreateAccountRequest struct {
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeat_password" binding:"required"`
}
