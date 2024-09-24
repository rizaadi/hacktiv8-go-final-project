package http

import (
	"hacktiv8-go-final-project/config"
	"hacktiv8-go-final-project/internal/domain/comment"
	"hacktiv8-go-final-project/internal/domain/photo"
	"hacktiv8-go-final-project/internal/domain/socialmedia"
	"hacktiv8-go-final-project/internal/domain/user"
	"hacktiv8-go-final-project/internal/infrastructure/db"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() (r *gin.Engine) {
	r = gin.New()

	//Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	dbGorm, err := config.InitDB()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// init repo & services
	userRepo := db.NewUserPostgresRepository(dbGorm)
	userService := user.NewService(userRepo)

	socialMediaRepo := db.NewSocialMediaPostgresRepository(dbGorm)
	socialMediaService := socialmedia.NewService(socialMediaRepo)

	photoRepo := db.NewPhotoPostgresRepository(dbGorm)
	photoService := photo.NewService(photoRepo)

	commentRepo := db.NewCommentPostgresRepository(dbGorm)
	commentService := comment.NewService(commentRepo)

	// pass services & *gin.Engin to init gin handler
	InitGinHandler(r, dbGorm, userService, socialMediaService, photoService, commentService)

	return
}
