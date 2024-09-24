package commenthandler

import (
	"hacktiv8-go-final-project/helpers"
	"hacktiv8-go-final-project/internal/domain/comment"
	"hacktiv8-go-final-project/internal/interfaces"
	commentDTO "hacktiv8-go-final-project/pkg/commentDTO"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service comment.Service
}

// Create Comment godoc
// @Summary      Create a comment
// @Description  Create a comment
// @Tags         comment
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        comment  body      commentDTO.Comment  true  "Create Comment"
// @Success      200   {object}  commentDTO.Comment
// @Router       /comment [post]
func (h *handler) Create(c *gin.Context) {
	var dto commentDTO.Comment
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

	comment := comment.Comment{
		UserID:  currentUser.ID,
		PhotoID: dto.PhotoID,
		Message: dto.Message,
	}

	if err := h.service.CreateComment(&comment); err != nil {
		log.Printf("Failed to create comment: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to create comment", err)
		return
	}

	helpers.RespondWithSuccess(c, "Comment created successfully", comment)
}

// Delete Comment godoc
// @Summary      Delete a comment
// @Description  Delete a comment
// @Tags         comment
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id  path  int  true  "Comment ID"
// @Success      200   {object}  string
// @Router       /comment/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Failed to parse id: %s", err)
		helpers.RespondIDInvalidError(c, err)
		return
	}

	if err := h.service.DeleteCommentByID(uint(id)); err != nil {
		log.Printf("Failed to delete comment: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to delete comment", err)
		return
	}

	helpers.RespondWithSuccess(c, "Comment deleted successfully", nil)
}

// GetAll Comment godoc
// @Summary      Get all comments
// @Description  Get all comments
// @Tags         comment
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        photo_id  query  int  false  "Photo ID"
// @Success      200   {object}  []commentDTO.CommentResponse
// @Router       /comment [get]
func (h *handler) GetAll(c *gin.Context) {
	currentUser := helpers.GetUserDataFromContext(c)
	comments, err := h.service.GetAllCommentsByPhotoID(currentUser.ID)

	if err != nil {
		log.Printf("Failed to get comments: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to get comments", err)
		return
	}

	var response []commentDTO.CommentResponse
	for _, comment := range comments {
		response = append(response, commentDTO.CommentResponse{
			ID:        comment.ID,
			PhotoID:   comment.PhotoID,
			Message:   comment.Message,
			CreatedAt: *comment.CreatedAt,
			UpdatedAt: *comment.UpdatedAt,
		})
	}

	helpers.RespondWithSuccess(c, "Comments retrieved successfully", response)
}

// GetByID Comment godoc
// @Summary      Get a comment by ID
// @Description  Get a comment by ID
// @Tags         comment
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id  path  int  true  "Comment ID"
// @Success      200   {object}  commentDTO.CommentResponse
// @Router       /comment/{id} [get]
func (h *handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("Failed to parse id: %s", err)
		helpers.RespondIDInvalidError(c, err)
		return
	}

	comment, err := h.service.GetCommentByID(uint(id))

	if err != nil {
		log.Printf("Failed to get comment: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to get comment", err)
		return
	}

	response := commentDTO.CommentResponse{
		ID:        comment.ID,
		PhotoID:   comment.PhotoID,
		Message:   comment.Message,
		CreatedAt: *comment.CreatedAt,
		UpdatedAt: *comment.UpdatedAt,
	}

	helpers.RespondWithSuccess(c, "Comment retrieved successfully", response)
}

// Update Comment godoc
// @Summary      Update a comment
// @Description  Update a comment
// @Tags         comment
// @Accept       json
// @Produce      json
// @Security 	 Bearer
// @Param        id  path  int  true  "Comment ID"
// @Param        comment  body      commentDTO.Comment  true  "Update Comment"
// @Success      200   {object}  commentDTO.CommentResponse
// @Router       /comment/{id} [put]
func (h *handler) Update(c *gin.Context) {
	var dto commentDTO.Comment
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

	comment := comment.Comment{
		GormModel: interfaces.GormModel{ID: uint(id)},
		Message:   dto.Message,
	}

	if err := h.service.UpdateComment(&comment); err != nil {
		log.Printf("Failed to update comment: %s", err)
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to update comment", err)
		return
	}

	response := commentDTO.CommentResponse{
		ID:        comment.ID,
		PhotoID:   comment.PhotoID,
		Message:   comment.Message,
		CreatedAt: *comment.CreatedAt,
		UpdatedAt: *comment.UpdatedAt,
	}

	helpers.RespondWithSuccess(c, "Comment updated successfully", response)
}

type Handler interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewCommentHandler(service comment.Service) Handler {
	return &handler{service}
}
