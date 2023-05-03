package handlers

import (
	"net/http"

	"github.com/gabehamasaki/finance-server/internal/dtos"
	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
)

func ShowUser(ctx *gin.Context) {
	id := ctx.Param("id")

	account := &models.Account{}

	err := db.First(&account, "id = ?", id).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "show-user", err.Error())
		return
	}

	response := &dtos.ShowAccountResponseDTO{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	}

	helpers.SendSuccess(ctx, "show-user", &response)
}
