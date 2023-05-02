package database

import (
	"os"

	"github.com/gabehamasaki/finance-server/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {

	var err error

	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Account{}, &models.Transaction{})

	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
