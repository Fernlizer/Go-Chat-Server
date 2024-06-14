// db/db.go
package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(localhost:3306)/gochat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", 0), // Use os.Stdout as the output, you can customize it
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = db
	// DB.AutoMigrate(
	// 	&models.User{},
	// 	&models.UserBlock{},
	// 	&models.Room{},
	// 	&models.RoomConfig{},
	// 	&models.RoomType{},
	// )
}
