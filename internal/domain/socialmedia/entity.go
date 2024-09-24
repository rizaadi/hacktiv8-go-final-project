package socialmedia

import (
	"hacktiv8-go-final-project/internal/domain/user"
	"hacktiv8-go-final-project/internal/interfaces"
)

type SocialMedia struct {
	interfaces.GormModel
	Name           string     `json:"name" gorm:"not null"`
	SocialMediaUrl string     `json:"social_media_url" gorm:"not null" `
	UserID         uint       `json:"user_id" gorm:"not null"`
	User           *user.User `json:"-"`
}
