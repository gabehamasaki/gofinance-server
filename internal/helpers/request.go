package helpers

import (
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAccount(ctx *gin.Context) *models.Account {
	var accountCTX interface{}

	accountCTX, _ = ctx.Get("account")

	account := accountCTX.(models.Account)

	return &account
}
