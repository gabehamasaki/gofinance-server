package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gabehamasaki/finance-server/internal/dtos"
	"github.com/gabehamasaki/finance-server/internal/helpers"
	"github.com/gabehamasaki/finance-server/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Auth(ctx *gin.Context) {
	request := &dtos.AuthRequest{}
	ctx.Bind(&request)

	if err := request.Validate(); err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "auth", err.Error())
		return
	}

	account := &models.Account{}

	if err := db.First(&account, "email = ?", request.Email).Error; err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "auth", "email or password is invalid")
		return
	}

	if !validatePassword(account.Password, request.Password) {
		helpers.SendError(ctx, http.StatusBadRequest, "auth", "email or password is invalid")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": account.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, "auth", err.Error())
		return
	}

	helpers.SendSuccess(ctx, "auth", gin.H{
		"token": tokenString,
	})
}

func validatePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
