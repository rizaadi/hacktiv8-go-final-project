package socialmediahandler

import (
	"hacktiv8-go-final-project/helpers"
	"hacktiv8-go-final-project/internal/domain/socialmedia"
	"hacktiv8-go-final-project/internal/interfaces"
	socialmediaDTO "hacktiv8-go-final-project/pkg/socialMediaDTO"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service socialmedia.Service
}

// Create Social Media godoc
// @Summary      Create a social media
// @Description  Create a social media
// @Tags         social media
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        social_media  body      socialmediaDTO.SocialMedia  true  "Create Social Media"
// @Success      200   {object}  socialmediaDTO.SocialMedia
// @Router       /social_media [post]
func (h *handler) Create(c *gin.Context) {
	var dto socialmediaDTO.SocialMedia
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

	socialMedia := socialmedia.SocialMedia{
		UserID:         currentUser.ID,
		Name:           dto.Name,
		SocialMediaUrl: dto.SocialMediaUrl,
	}

	if err := h.service.CreateSocialMedia(&socialMedia); err != nil {
		log.Printf("Failed to create social media: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to create social media", err)
		return
	}

	helpers.RespondWithSuccess(c, "Social media created successfully", socialMedia)
}

// Delete Social Media godoc
// @Summary      Delete a social media
// @Description  Delete a social media
// @Tags         social media
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id  path  int  true  "Social Media ID"
// @Success      200   {object}  string
// @Router       /social_media/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Failed to parse id: %s", err)
		helpers.RespondIDInvalidError(c, err)
		return
	}

	if err := h.service.DeleteSocialMediaByID(uint(id)); err != nil {
		log.Printf("Failed to delete social media: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to delete social media", err)
		return
	}

	helpers.RespondWithSuccess(c, "Social media deleted successfully", nil)
}

// GetAll Social Media godoc
// @Summary      Get all social media
// @Description  Get all social media
// @Tags         social media
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Success      200   {object}  []socialmediaDTO.SocialMediaResponse
// @Router       /social_media [get]
func (h *handler) GetAll(c *gin.Context) {
	currentUser := helpers.GetUserDataFromContext(c)
	socialMedias, err := h.service.GetAllSocialMediasByUserID(currentUser.ID)

	if err != nil {
		log.Printf("Failed to get social media: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to get social media", err)
		return
	}

	var response []socialmediaDTO.SocialMediaResponse
	for _, socialMedia := range socialMedias {
		response = append(response, socialmediaDTO.SocialMediaResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			CreatedAt:      *socialMedia.CreatedAt,
			UpdatedAt:      *socialMedia.UpdatedAt,
		})
	}

	helpers.RespondWithSuccess(c, "Social media retrieved successfully", response)
}

// GetByID Social Media godoc
// @Summary      Get a social media by ID
// @Description  Get a social media by ID
// @Tags         social media
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id  path  int  true  "Social Media ID"
// @Success      200   {object}  socialmediaDTO.SocialMediaResponse
// @Router       /social_media/{id} [get]
func (h *handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Failed to parse id: %s", err)
		helpers.RespondIDInvalidError(c, err)
		return
	}

	socialMedia, err := h.service.GetSocialMediaByID(uint(id))

	if err != nil {
		log.Printf("Failed to get social media: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to get social media", err)
		return
	}

	response := socialmediaDTO.SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		CreatedAt:      *socialMedia.CreatedAt,
		UpdatedAt:      *socialMedia.UpdatedAt,
	}

	helpers.RespondWithSuccess(c, "Social media retrieved successfully", response)
}

// Update Social Media godoc
// @Summary      Update a social media
// @Description  Update a social media
// @Tags         social media
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id  path  int  true  "Social Media ID"
// @Param        social_media  body  socialmediaDTO.UpdateSocialMedia  true  "Update Social Media"
// @Success      200   {object}  socialmediaDTO.SocialMediaResponse
// @Router       /social_media/{id} [put]
func (h *handler) Update(c *gin.Context) {
	var dto socialmediaDTO.UpdateSocialMedia
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

	socialMedia := socialmedia.SocialMedia{
		GormModel:      interfaces.GormModel{ID: uint(id)},
		Name:           dto.Name,
		SocialMediaUrl: dto.SocialMediaUrl,
	}

	if err := h.service.UpdateSocialMedia(&socialMedia); err != nil {
		log.Printf("Failed to update social media: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to update social media", err)
		return
	}

	response := socialmediaDTO.SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		CreatedAt:      *socialMedia.CreatedAt,
		UpdatedAt:      *socialMedia.UpdatedAt,
	}

	helpers.RespondWithSuccess(c, "Social media updated successfully", response)
}

type Handler interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewSocialMediaHandler(s socialmedia.Service) Handler {
	return &handler{service: s}
}
