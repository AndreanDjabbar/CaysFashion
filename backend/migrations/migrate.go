package migrations

import (
	"github.com/AndreanDjabbar/CaysFashion/backend/pkg/logger"
	"gorm.io/gorm"
)

var log = logger.SetUpLogger()

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(

	); err != nil {
		log.Error("Failed to migrate", "error", err)
		panic(err)
	}
}