package socialmediaDTO

import (
	"github.com/asaskevich/govalidator"
	"time"
)

// SocialMedia represents the social media model.
type SocialMedia struct {
	Name           string `json:"name" valid:"required~Name is required, minstringlength(4)~Name has to have a minimum length of 4 characters" example:"Facebook"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Social Media Url is required" example:"https://www.facebook.com"`
}

// UpdateSocialMedia represents the update social media model.
type UpdateSocialMedia struct {
	SocialMedia
}

// SocialMediaResponse represents the social media response model.
type SocialMediaResponse struct {
	ID             uint      `json:"id" example:"1"`
	Name           string    `json:"name" example:"Facebook"`
	SocialMediaUrl string    `json:"social_media_url" example:"https://www.facebook.com"`
	CreatedAt      time.Time `json:"created_at" example:"2021-08-01T00:00:00Z"`
	UpdatedAt      time.Time `json:"updated_at" example:"2021-08-01T00:00:00Z"`
}

func (dto *SocialMedia) Validate() error {
	_, err := govalidator.ValidateStruct(dto)
	return err
}

func (dto *UpdateSocialMedia) Validate() error {
	_, err := govalidator.ValidateStruct(dto)
	return err
}
