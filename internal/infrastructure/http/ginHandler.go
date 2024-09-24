package http

import (
	"hacktiv8-go-final-project/internal/domain/comment"
	"hacktiv8-go-final-project/internal/domain/photo"
	"hacktiv8-go-final-project/internal/domain/socialmedia"
	"hacktiv8-go-final-project/internal/domain/user"
	commentHandler "hacktiv8-go-final-project/internal/interfaces/commentHandler"
	"hacktiv8-go-final-project/internal/interfaces/middleware"
	photohandler "hacktiv8-go-final-project/internal/interfaces/photoHandler"
	socialMediaHandler "hacktiv8-go-final-project/internal/interfaces/socialMediaHandler"
	userHandler "hacktiv8-go-final-project/internal/interfaces/userHandler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "hacktiv8-go-final-project/docs"
)

// @title MyGram API
// @version 1.0
// @description This is a simple MyGram API
// @termsOfService http://swagger.io/terms/
// @contact.name Riza
// @contact.url http://www.rizaadikurniawan.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @SecurityDefinitions.apiKey Bearer
// @in header
// @name Authorization
func InitGinHandler(r *gin.Engine, db *gorm.DB, userService user.Service, socialMediaService socialmedia.Service, photoService photo.Service, commentService comment.Service) {
	userhandler := userHandler.NewUserHandler(userService)
	socialMediahandler := socialMediaHandler.NewSocialMediaHandler(socialMediaService)
	photohandler := photohandler.NewPhotoHandler(photoService)
	commentHandler := commentHandler.NewCommentHandler(commentService)

	r.POST("/register", userhandler.CreateUser)
	r.POST("/login", userhandler.Login)

	socialMediaRouter := r.Group("/social_media")
	socialMediaRouter.Use(middleware.Authentication())
	socialMediaRouter.POST("/", socialMediahandler.Create)
	socialMediaRouter.GET("/", socialMediahandler.GetAll)
	socialMediaRouter.GET("/:id", socialMediahandler.GetByID)
	socialMediaRouter.Use(middleware.SocialMediaAuthorization(db))
	socialMediaRouter.PUT("/:id", middleware.SocialMediaAuthorization(db), socialMediahandler.Update)
	socialMediaRouter.DELETE("/:id", middleware.SocialMediaAuthorization(db), socialMediahandler.Delete)

	photoRouter := r.Group("/photo")
	photoRouter.Use(middleware.Authentication())
	photoRouter.POST("/", photohandler.Create)
	photoRouter.GET("/", photohandler.GetAll)
	photoRouter.GET("/:id", photohandler.GetByID)
	photoRouter.Use(middleware.PhotoAuthorization(db))
	photoRouter.PUT("/:id", middleware.PhotoAuthorization(db), photohandler.Update)
	photoRouter.DELETE("/:id", middleware.PhotoAuthorization(db), photohandler.Delete)

	commentRouter := r.Group("/comment")
	commentRouter.Use(middleware.Authentication())
	commentRouter.POST("/", commentHandler.Create)
	commentRouter.GET("/", commentHandler.GetAll)
	commentRouter.GET("/:id", commentHandler.GetByID)
	commentRouter.Use(middleware.CommentAuthorization(db))
	commentRouter.PUT("/:id", middleware.CommentAuthorization(db), commentHandler.Update)
	commentRouter.DELETE("/:id", middleware.CommentAuthorization(db), commentHandler.Delete)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
