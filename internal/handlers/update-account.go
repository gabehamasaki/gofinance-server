package handlers

import (
	"net/http"

	"github.com/gabehamasaki/finance-server/internal/dtos"
	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gin-gonic/gin"
)

func UpdateAccount(ctx *gin.Context) {
	account := helpers.GetAccount(ctx)

	request := &dtos.UpdateAccountDTO{}
	ctx.Bind(&request)

	if err := request.Validate(); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "update-account", err.Error())
		return
	}

	if request.Name != "" {
		account.Name = request.Name
	}

	if request.Email != "" {
		account.Email = request.Email
	}

	if request.Password != "" {
		account.Password = encryptPassword(request.Password)
	}

	if err := db.Save(&account).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "update-account", err.Error())
		return
	}

	response := &dtos.UpdateAccountResponseDTO{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	}

	helpers.SendSuccess(ctx, "update-account", &response)
}
