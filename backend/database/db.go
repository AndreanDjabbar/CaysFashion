package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/AndreanDjabbar/CaysFashion/backend/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
    db   *gorm.DB
    once sync.Once
)

func GetDB() *gorm.DB {
    logger := logger.SetUpLogger()

    once.Do(func() {
        var err error
        dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
        )
        fmt.Println(dsn)
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
            logger.Error(
                "Failed to connect to database",
                "error", err,
            )
        }
    })
    return db
}
