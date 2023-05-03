package handlers

import (
	"net/http"

	"github.com/gabehamasaki/finance-server/internal/dtos"
	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
)

func UpdateTransactio(ctx *gin.Context) {
	id := ctx.Param("id")

	transaction := &models.Transaction{}
	err := db.First(&transaction, "id = ?", id).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "update-transaction", err.Error())
		return
	}

	request := &dtos.UpdateTransactionRequestDTO{}
	ctx.Bind(&request)

	if err := request.Validate(); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "update-transaction", err.Error())
		return
	}

	if request.Title != "" {
		transaction.Title = request.Title
	}

	if request.Description != nil {
		transaction.Description = request.Description
	}

	if request.Type != "" {
		transaction.Type = request.Type
	}

	if request.Value <= 0 {
		transaction.Value = request.Value
	}

	if err := db.Save(&transaction).Error; err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "update-transaction", err.Error())
		return
	}

	response := &dtos.UpdateTransactionResponseDTO{
		ID:          transaction.ID,
		Title:       transaction.Title,
		Description: transaction.Description,
		Value:       transaction.Value,
		Type:        transaction.Type,
	}

	helpers.SendSuccess(ctx, "update-transaction", &response)
}
