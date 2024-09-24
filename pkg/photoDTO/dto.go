package photoDTO

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Photo struct {
	Title    string `json:"title" valid:"required~Title is required, minstringlength(4)~Title has to have a minimum length of 4 characters" example:"My Photo"`
	Caption  string `json:"caption" valid:"required~Caption is required" example:"This is my photo"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo URL is required" example:"https://www.google.com/image.jpg"`
}

type UpdatePhoto struct {
	Photo
}
type PhotoResponse struct {
	ID        uint      `json:"id" example:"1"`
	Title     string    `json:"title" example:"My Photo"`
	Caption   string    `json:"caption" example:"This is my photo"`
	PhotoUrl  string    `json:"photo_url" example:"https://www.google.com/image.jpg"`
	CreatedAt time.Time `json:"created_at" example:"2021-07-12T09:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-07-12T09:00:00Z"`
}

func (dto *Photo) Validate() error {
	_, err := govalidator.ValidateStruct(dto)
	return err
}

func (dto *UpdatePhoto) Validate() error {
	_, err := govalidator.ValidateStruct(dto)
	return err
}
