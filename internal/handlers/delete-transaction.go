package handlers

import (
	"net/http"

	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
)

func DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")

	transaction := &models.Transaction{}

	err := db.First(&transaction, "id = ?", id).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "delete-transaction", err.Error())
		return
	}

	err = db.Delete(&transaction).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "delete-transaction", err.Error())
		return
	}

	helpers.SendSuccess(ctx, "delete-transaction", &transaction)
}
