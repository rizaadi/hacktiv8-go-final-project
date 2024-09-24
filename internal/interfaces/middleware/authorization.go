package middleware

import (
	"hacktiv8-go-final-project/helpers"
	"hacktiv8-go-final-project/internal/domain/comment"
	"hacktiv8-go-final-project/internal/domain/photo"
	"hacktiv8-go-final-project/internal/domain/socialmedia"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SocialMediaAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		socialMediaID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			helpers.RespondBindJSONError(c, err)
			return
		}

		userData := helpers.GetUserDataFromContext(c)

		var socialMedia socialmedia.SocialMedia
		if err := db.First(&socialMedia, socialMediaID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				helpers.RespondWithError(c, http.StatusNotFound, "Social media not found", err)
				return
			}
			helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to get social media", err)
			return
		}

		if socialMedia.UserID != userData.ID {
			helpers.RespondWithError(c, http.StatusForbidden, "You do not have permission to access this resource", nil)
			return
		}

		c.Next()
	}
}

func PhotoAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		photoID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			helpers.RespondBindJSONError(c, err)
			return
		}

		userData := helpers.GetUserDataFromContext(c)

		var photo photo.Photo
		if err := db.First(&photo, photoID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				helpers.RespondWithError(c, http.StatusNotFound, "Photo not found", err)
				return
			}
			helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to get photo", err)
			return
		}

		if photo.UserID != userData.ID {
			helpers.RespondWithError(c, http.StatusForbidden, "You do not have permission to access this resource", nil)
			return
		}

		c.Next()
	}
}

func CommentAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			helpers.RespondBindJSONError(c, err)
			return
		}

		userData := helpers.GetUserDataFromContext(c)

		var comment comment.Comment
		if err := db.First(&comment, commentID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				helpers.RespondWithError(c, http.StatusNotFound, "Comment not found", err)
				return
			}
			helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to get comment", err)
			return
		}

		if comment.UserID != userData.ID {
			helpers.RespondWithError(c, http.StatusForbidden, "You do not have permission to access this resource", nil)
			return
		}

		c.Next()
	}
}
