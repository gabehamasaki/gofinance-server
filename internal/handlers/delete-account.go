package handlers

import (
	"net/http"

	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
)

func DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")

	account := &models.Account{}

	err := db.First(&account, "id = ?", id).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "delete-account", err.Error())
		return
	}

	transactions := &[]models.Transaction{}

	err = db.Find(&transactions, "owner = ?", id).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "delete-account", err.Error())
		return
	}

	err = db.Delete(&transactions, "owner = ?", id).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "delete-account", err.Error())
		return
	}

	err = db.Delete(&account).Error
	if err != nil {
		helpers.SendError(ctx, http.StatusInternalServerError, "delete-account", err.Error())
		return
	}

	helpers.SendSuccess(ctx, "delete-account", &account)
}
