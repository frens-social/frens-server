package database

import (
	"log"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

func Connect() {
	dbURL := "postgres://" +
		config.C.Database.User +
		":" + config.C.Database.Password +
		"@" + config.C.Database.Host +
		":" + config.C.Database.Port +
		"/" + config.C.Database.Database

	var err error
	database, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
		//Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalln(err)
	}

	if err := database.AutoMigrate(
		&models.Account{},
	); err != nil {
		log.Fatalln(err)
	}
}
