package handlers

import (
	"net/http"

	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
)

func ShowTransaction(ctx *gin.Context) {
	id := ctx.Param("id")

	transaction := &models.Transaction{}

	err := db.First(&transaction, "id = ?", id).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "show-transaction", err.Error())
		return
	}

	helpers.SendSuccess(ctx, "show-transaction", &transaction)
}
