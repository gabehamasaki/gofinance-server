package helpers

import "github.com/gin-gonic/gin"

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"op":   op,
		"data": data,
	})
}

func SendError(ctx *gin.Context, code int, op string, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"errorCode": code,
		"op":        op,
		"message ":  msg,
	})
}

func SendErrorData(ctx *gin.Context, code int, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"errorCode": code,
		"op":        op,
		"data ":     data,
	})
}
