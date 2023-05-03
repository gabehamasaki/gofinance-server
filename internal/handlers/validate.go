package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateToken(ctx *gin.Context) {
	authenticationHeader := ctx.GetHeader("Authorization")

	if authenticationHeader == "" {
		helpers.SendError(ctx, http.StatusBadRequest, "validate-auth", "Invalid token")
		return
	}

	tokenString := strings.Split(authenticationHeader, " ")[1]

	if tokenString == "" {
		helpers.SendError(ctx, http.StatusBadRequest, "validate-auth", "Invalid token")
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header)
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		helpers.SendErrorData(ctx, http.StatusBadRequest, "validate-auth", gin.H{
			"valid": false,
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		helpers.SendErrorData(ctx, http.StatusBadRequest, "validate-auth", gin.H{
			"valid": false,
		})
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		helpers.SendErrorData(ctx, http.StatusBadRequest, "validate-auth", gin.H{
			"valid": false,
		})
		return
	}

	var account models.Account
	if err := db.First(&account, "id = ?", claims["sub"]).Error; err != nil {
		helpers.SendErrorData(ctx, http.StatusBadRequest, "validate-auth", gin.H{
			"valid": false,
		})
		return
	}

	helpers.SendSuccess(ctx, "validate-auth", gin.H{
		"valid": true,
	})
}
