package commentDTO

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Comment represents the comment model.
type Comment struct {
	PhotoID uint   `json:"photo_id" valid:"required~Photo ID is required" example:"1"`
	Message string `json:"message" valid:"required~Comment is required" example:"This is my comment"`
}

// UpdateComment represents the update comment model.
type UpdateComment struct {
	Comment
}

// CommentResponse represents the comment response model.
type CommentResponse struct {
	ID        uint      `json:"id" example:"1"`
	PhotoID   uint      `json:"photo_id" example:"1"`
	Message   string    `json:"message" example:"This is my comment"`
	CreatedAt time.Time `json:"created_at" example:"2021-07-12T09:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-07-12T09:00:00Z"`
}

func (dto *Comment) Validate() error {
	_, err := govalidator.ValidateStruct(dto)
	return err
}

func (dto *UpdateComment) Validate() error {
	_, err := govalidator.ValidateStruct(dto)
	return err
}
