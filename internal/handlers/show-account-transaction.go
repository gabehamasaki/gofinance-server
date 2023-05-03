package handlers

import (
	"net/http"

	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
)

func ShowAccountTransactions(ctx *gin.Context) {
	id := helpers.GetAccount(ctx).ID.String()

	types, exist := ctx.GetQuery("type")

	if exist && types != "" {
		transactions := &[]models.Transaction{}

		err := db.Find(&transactions, "owner = ? and type = ?", id, types).Error
		if err != nil {
			helpers.SendError(ctx, http.StatusBadRequest, "show-account-transactions", err.Error())
			return
		}

		helpers.SendSuccess(ctx, "show-account-transactions", &transactions)
		return
	}

	transactions := &[]models.Transaction{}

	err := db.Find(&transactions, "owner = ?", id).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "show-account-transactions", err.Error())
		return
	}

	helpers.SendSuccess(ctx, "show-account-transactions", &transactions)
}
