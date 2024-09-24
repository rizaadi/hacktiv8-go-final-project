package userhandler

import (
	"errors"
	"hacktiv8-go-final-project/helpers"
	"hacktiv8-go-final-project/internal/domain/user"

	userDTO "hacktiv8-go-final-project/pkg/userDTO"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	service user.Service
}

// Login godoc
// @Summary      Login a user
// @Description  Authenticate a user and return a JWT token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      userDTO.LoginUserRequest  true  "Login User"
// @Success      200   {object}  userDTO.LoginUserResponse
// @Router       /login [post]
func (h *handler) Login(c *gin.Context) {
	var param userDTO.LoginUserRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		log.Printf("Failed to bind request JSON: %s", err)
		helpers.RespondBindJSONError(c, err)
		return
	}

	userData, err := h.service.GetUserByEmail(param.Email)
	if err != nil {
		log.Printf("Failed get user: %s", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.RespondWithError(c, http.StatusBadRequest, helpers.ErrAccountNotFound.Error(), err)
		}
		return

	}

	if !helpers.ComparePassword([]byte(userData.Password), []byte(param.Password)) {
		helpers.RespondWithError(c, http.StatusBadRequest, "Wrong Password", err)
		return
	}
	token, err := helpers.GenerateJWT(helpers.Claims{
		ID:    userData.ID,
		Email: userData.Email,
	})
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	helpers.RespondWithSuccess(c, "Login success", userDTO.LoginUserResponse{
		Username: userData.UserName,
		Age:      userData.Age,
		Email:    userData.Email,
		Token:    token,
	})
}

// CreateUser godoc
// @Summary      Create a user
// @Description  Create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      userDTO.User  true  "Create User"
// @Success      200   {object}  userDTO.User
// @Router       /register [post]
func (h *handler) CreateUser(c *gin.Context) {
	var param userDTO.User

	if err := c.ShouldBindJSON(&param); err != nil {
		log.Printf("Failed to bind request JSON: %s", err)
		helpers.RespondBindJSONError(c, err)
		return
	}
	if err := param.Validate(); err != nil {
		helpers.RespondValidatorError(c, err)
		return
	}

	userData := user.User{
		UserName: param.Username,
		Age:      param.Age,
		Email:    param.Email,
		Password: param.Password,
	}

	if err := h.service.CreateUser(&userData); err != nil {
		log.Printf("Failed create user: %s", err)
		if err == helpers.ErrEmailAlreadyRegistered {
			helpers.RespondWithError(c, http.StatusInternalServerError, err.Error(), err)
		} else if err == helpers.ErrAgeMustBeMore {
			helpers.RespondWithError(c, http.StatusInternalServerError, err.Error(), err)
		} else {
			helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to create user", err)
		}
		return
	}

	helpers.RespondWithSuccess(c, "User created successfully", userData)
}

type Handler interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
}

func NewUserHandler(s user.Service) Handler {
	return &handler{service: s}
}
