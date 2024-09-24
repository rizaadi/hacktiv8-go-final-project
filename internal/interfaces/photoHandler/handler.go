package photohandler

import (
	"hacktiv8-go-final-project/helpers"
	"hacktiv8-go-final-project/internal/domain/photo"
	"hacktiv8-go-final-project/internal/interfaces"
	photoDTO "hacktiv8-go-final-project/pkg/photoDTO"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service photo.Service
}

// Create Photo godoc
// @Summary      Create a photo
// @Description  Create a photo
// @Tags         photo
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        photo  body      photoDTO.Photo  true  "Create Photo"
// @Success      200   {object}  photoDTO.Photo
// @Router       /photo [post]
func (h *handler) Create(c *gin.Context) {
	var dto photoDTO.Photo
	currentUser := helpers.GetUserDataFromContext(c)

	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("Failed to bind request JSON: %s", err)
		helpers.RespondBindJSONError(c, err)
		return
	}

	if err := dto.Validate(); err != nil {
		helpers.RespondValidatorError(c, err)
		return
	}

	photo := photo.Photo{
		UserID:   currentUser.ID,
		Title:    dto.Title,
		Caption:  dto.Caption,
		PhotoUrl: dto.PhotoUrl,
	}

	if err := h.service.CreatePhoto(&photo); err != nil {
		log.Printf("Failed to create photo: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to create photo", err)
		return
	}

	helpers.RespondWithSuccess(c, "Photo created successfully", photo)

}

// Delete Photo godoc
// @Summary      Delete a photo
// @Description  Delete a photo
// @Tags         photo
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id  path  int  true  "Photo ID"
// @Success      200   {object}  string
// @Router       /photo/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Failed to parse id: %s", err)
		helpers.RespondIDInvalidError(c, err)
		return
	}

	if err := h.service.DeletePhotoByID(uint(id)); err != nil {
		log.Printf("Failed to delete photo: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to delete photo", err)
		return
	}

	helpers.RespondWithSuccess(c, "Photo deleted successfully", nil)
}

// GetAll Photo godoc
// @Summary      Get all photos
// @Description  Get all photos
// @Tags         photo
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Success      200   {object}  []photoDTO.PhotoResponse
// @Router       /photo [get]
func (h *handler) GetAll(c *gin.Context) {
	currentUser := helpers.GetUserDataFromContext(c)
	photos, err := h.service.GetAllPhotosByUserID(currentUser.ID)

	if err != nil {
		log.Printf("Failed to get photos: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to get photos", err)
		return
	}
	var response []photoDTO.PhotoResponse
	for _, photo := range photos {
		response = append(response, photoDTO.PhotoResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			CreatedAt: *photo.CreatedAt,
			UpdatedAt: *photo.UpdatedAt,
		})
	}

	helpers.RespondWithSuccess(c, "Photos retrieved successfully", response)
}

// GetByID Photo godoc
// @Summary      Get a photo by ID
// @Description  Get a photo by ID
// @Tags         photo
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id  path  int  true  "Photo ID"
// @Success      200   {object}  photoDTO.PhotoResponse
// @Router       /photo/{id} [get]
func (h *handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Failed to parse id: %s", err)
		helpers.RespondIDInvalidError(c, err)
		return
	}

	photo, err := h.service.GetPhotoByID(uint(id))

	if err != nil {
		log.Printf("Failed to get photo: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to get photo", err)
		return
	}

	response := photoDTO.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		CreatedAt: *photo.CreatedAt,
		UpdatedAt: *photo.UpdatedAt,
	}

	helpers.RespondWithSuccess(c, "Photo retrieved successfully", response)
}

// Update Photo godoc
// @Summary      Update a photo
// @Description  Update a photo
// @Tags         photo
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id  path  int  true  "Photo ID"
// @Param        photo  body      photoDTO.UpdatePhoto  true  "Update Photo"
// @Success      200   {object}  photoDTO.PhotoResponse
// @Router       /photo/{id} [put]
func (h *handler) Update(c *gin.Context) {
	var dto photoDTO.UpdatePhoto
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Failed to parse id: %s", err)
		helpers.RespondIDInvalidError(c, err)
		return
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("Failed to bind request JSON: %s", err)
		helpers.RespondBindJSONError(c, err)
		return
	}

	if err := dto.Validate(); err != nil {
		helpers.RespondValidatorError(c, err)
		return
	}

	photo := photo.Photo{
		GormModel: interfaces.GormModel{ID: uint(id)},
		Title:     dto.Title,
		Caption:   dto.Caption,
		PhotoUrl:  dto.PhotoUrl,
	}

	if err := h.service.UpdatePhoto(&photo); err != nil {
		log.Printf("Failed to update photo: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to update photo", err)
		return
	}

	response := photoDTO.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		CreatedAt: *photo.CreatedAt,
		UpdatedAt: *photo.UpdatedAt,
	}

	helpers.RespondWithSuccess(c, "Photo updated successfully", response)
}

type Handler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
}

func NewPhotoHandler(service photo.Service) Handler {
	return &handler{service}
}
