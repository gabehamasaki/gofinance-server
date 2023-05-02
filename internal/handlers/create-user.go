package handlers

import (
	"net/http"
	"strings"

	"github.com/gabehamasaki/finance-server/internal/dtos"
	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx *gin.Context) {
	request := dtos.CreateAccountDTO{}

	ctx.Bind(&request)

	if err := request.Validate(); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "create-account", err.Error())
		return
	}

	account := &models.Account{
		Name:     request.Name,
		Email:    request.Email,
		Password: encryptPassword(request.Password),
	}
	account.GenerateId()

	if err := db.Create(account).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			helpers.SendError(ctx, http.StatusFound, "create-account", "email already exists")
			return
		}
		helpers.SendError(ctx, http.StatusInternalServerError, "create-account", err.Error())
		return
	}

	helpers.SendSuccess(ctx, "create-account", &account)
}

func encryptPassword(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 8)

	return string(hash)
}
