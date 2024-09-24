package user

import (
	"hacktiv8-go-final-project/helpers"
	"hacktiv8-go-final-project/internal/interfaces"

	"gorm.io/gorm"
)

type User struct {
	interfaces.GormModel
	UserName string `json:"user_name" gorm:"not null;uniqueIndex"`
	Age      int    `json:"age" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;uniqueIndex"`
	Password string `json:"password" gorm:"not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var count int64

	// Check if email is already registered
	if err := tx.Model(&User{}).Where("email = ?", u.Email).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return helpers.ErrEmailAlreadyRegistered
	}

	// Reset count and check if username is already registered
	count = 0
	if err := tx.Model(&User{}).Where("user_name = ?", u.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return helpers.ErrUserNameAlreadyRegistered
	}

	// Check if age is valid
	if u.Age < 8 {
		return helpers.ErrAgeMustBeMore
	}

	// Hash the password
	u.Password = helpers.HashPassword(u.Password)
	return nil
}
