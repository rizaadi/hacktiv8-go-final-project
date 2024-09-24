package userDTO

import "github.com/asaskevich/govalidator"

// User represents the user model.
type User struct {
	Username string `json:"user_name" valid:"required~Username name is required, minstringlength(4)~Username has to have a minimum length of 4 characters" example:"riza"`
	Age      int    `json:"age" valid:"required~Your Age is required" example:"18"`
	Email    string `json:"email" valid:"email,required~Your Email is required" example:"riza@gmail.com"`
	Password string `json:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters" example:"rza123"`
}

// LoginUserRequest represents the login user model.
type LoginUserRequest struct {
	Email    string `json:"email" valid:"email,required~Your Email is required" example:"riza@gmail.com"`
	Password string `json:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters" example:"rza123"`
}

type LoginUserResponse struct {
	Username string `json:"user_name" example:"riza"`
	Age      int    `json:"age" example:"18"`
	Email    string `json:"email" example:"riza@gmail.com"`
	Token    string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjYwNzQwNzcsImlkIjoxLCJlbWFpbCI6InJpemFAZ21haWwuY29tIn0.7"`
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

func (lu *LoginUserRequest) Validate() error {
	_, err := govalidator.ValidateStruct(lu)
	return err
}
