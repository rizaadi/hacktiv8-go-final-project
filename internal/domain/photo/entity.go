package photo

import (
	"hacktiv8-go-final-project/internal/domain/user"
	"hacktiv8-go-final-project/internal/interfaces"
)

type Photo struct {
	interfaces.GormModel
	Title    string     `json:"title" gorm:"not null"`
	Caption  string     `json:"caption" gorm:"not null"`
	PhotoUrl string     `json:"photo_url" gorm:"not null"`
	UserID   uint       `json:"user_id" gorm:"not null"`
	User     *user.User `json:"-"`
}
