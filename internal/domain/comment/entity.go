package comment

import (
	"hacktiv8-go-final-project/internal/domain/photo"
	"hacktiv8-go-final-project/internal/domain/user"
	"hacktiv8-go-final-project/internal/interfaces"
)

type Comment struct {
	interfaces.GormModel
	Message string       `json:"message" gorm:"not null"`
	PhotoID uint         `json:"photo_id" gorm:"not null"`
	Photo   *photo.Photo `json:"-"`
	UserID  uint         `json:"user_id" gorm:"not null"`
	User    *user.User   `json:"-"`
}
