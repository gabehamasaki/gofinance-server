package handlers

import (
	"github.com/gabehamasaki/finance-server/internal/dtos"
	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gin-gonic/gin"
)

func ShowUser(ctx *gin.Context) {

	account := helpers.GetAccount(ctx)

	response := &dtos.ShowAccountResponseDTO{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	}

	helpers.SendSuccess(ctx, "show-user", &response)
}
