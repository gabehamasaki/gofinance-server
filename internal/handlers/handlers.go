package handlers

import (
	"github.com/gabehamasaki/finance-server/internal/database"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() {
	db = database.GetDatabase()
}
