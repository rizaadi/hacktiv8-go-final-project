package config

import (
	"fmt"
	"hacktiv8-go-final-project/internal/domain/comment"
	"hacktiv8-go-final-project/internal/domain/photo"
	"hacktiv8-go-final-project/internal/domain/socialmedia"
	"hacktiv8-go-final-project/internal/domain/user"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		err = fmt.Errorf("failed to connect to database: %v", err)
		return
	}
	if err = db.AutoMigrate(&user.User{}, &socialmedia.SocialMedia{}, &photo.Photo{}, &comment.Comment{}); err != nil {
		err = fmt.Errorf("failed to run auto migrations: %s", err)
	} else {
		log.Println("Auto migration completed")
	}
	return
}
