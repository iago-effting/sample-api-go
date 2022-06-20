package handlers

import (
	"fmt"
	"iago-effting/api-example/pkg/logs"
	"net/http"

	"github.com/gin-gonic/gin"

	"iago-effting/api-example/pkg/accounts"
	accountDb "iago-effting/api-example/pkg/storage/database/account"
)

func IndexAccount(ctx *gin.Context) {
	var usersRepository accounts.Repository = accountDb.Repo()
	var data *[]accounts.User

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
	var usersRepository accounts.Repository = accountDb.Repo()
	var identity = ctx.MustGet("identity").(string)
	var data *accounts.User

	data, err := usersRepository.GetBy(ctx, "email", identity)

	logger.Info(fmt.Sprintf("Getting view accounts for ID %s", data.ID))

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
	var usersRepository accounts.Repository = accountDb.Repo()

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
	var usersRepository accounts.Repository = accountDb.Repo()

	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := accounts.User{
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	newUser, err := usersRepository.Save(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
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
