package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gabehamasaki/finance-server/internal/database"
	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(ctx *gin.Context) {
	tokenString := strings.Split(ctx.GetHeader("Authorization"), " ")[1]

	if tokenString == "" {
		helpers.SendError(ctx, http.StatusUnauthorized, "auth-middleware", "Invalid token")
		ctx.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header)
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		helpers.SendError(ctx, http.StatusUnauthorized, "auth-middleware", err.Error())
		ctx.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			helpers.SendError(ctx, http.StatusUnauthorized, "auth-middleware", "Invalid token")
			ctx.Abort()
			return
		}

		var account models.Account
		if err := database.GetDatabase().First(&account, "id = ?", claims["sub"]).Error; err != nil {
			helpers.SendError(ctx, http.StatusUnauthorized, "auth-middleware", "Invalid token")
			ctx.Abort()
			return
		}

		ctx.Set("account", &account)

		ctx.Next()
	} else {
		helpers.SendError(ctx, http.StatusUnauthorized, "auth-middleware", "Invalid token")
		ctx.Abort()
		return
	}
}
