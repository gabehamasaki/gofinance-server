package handlers

import (
	"net/http"

	"github.com/gabehamasaki/finance-server/internal/dtos"
	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(ctx *gin.Context) {
	request := &dtos.CreateTransactionDTO{}

	ctx.Bind(&request)

	if err := request.Validate(); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "create-transaction", err.Error())
		return
	}

	var existOwner *models.Account

	err := db.First(&existOwner, "id = ?", request.Owner).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "create-transaction", "owner not exists")
		return
	}

	transaction := &models.Transaction{
		Title:       request.Title,
		Description: request.Description,
		Value:       request.Value,
		Type:        request.Type,
		Owner:       request.Owner,
	}
	transaction.GenerateId()

	err = db.Create(&transaction).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "create-transaction", err.Error())
		return
	}

	response := &dtos.TransactionResponseDTO{
		ID:          transaction.ID,
		Title:       transaction.Title,
		Description: transaction.Description,
		Value:       transaction.Value,
		Type:        transaction.Type,
	}

	helpers.SendSuccess(ctx, "create-transaction", &response)
}
