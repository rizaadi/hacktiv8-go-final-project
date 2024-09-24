package db

import (
	"hacktiv8-go-final-project/internal/domain/comment"
	"hacktiv8-go-final-project/internal/domain/photo"
	"hacktiv8-go-final-project/internal/domain/socialmedia"
	"hacktiv8-go-final-project/internal/domain/user"

	"gorm.io/gorm"
)

type pgRepository struct {
	db *gorm.DB
}

// CreateComment implements comment.CommentRepository.
func (p *pgRepository) CreateComment(comment *comment.Comment) (err error) {
	return p.db.Create(comment).Error
}

// DeleteCommentByID implements comment.CommentRepository.
func (p *pgRepository) DeleteCommentByID(id uint) (err error) {
	return p.db.Delete(&comment.Comment{}, id).Error
}

// GetAllCommentsByPhotoID implements comment.CommentRepository.
func (p *pgRepository) GetAllCommentsByPhotoID(photoID uint) (comments []comment.Comment, err error) {
	return comments, p.db.Where(&comment.Comment{PhotoID: photoID}).Find(&comments).Error
}

// GetCommentByID implements comment.CommentRepository.
func (p *pgRepository) GetCommentByID(id uint) (comment comment.Comment, err error) {
	return comment, p.db.First(&comment, id).Error
}

// UpdateComment implements comment.CommentRepository.
func (p *pgRepository) UpdateComment(comment *comment.Comment) (err error) {
	return p.db.Updates(comment).Error
}

// CreatePhoto implements photo.PhotoRepository.
func (p *pgRepository) CreatePhoto(photo *photo.Photo) (err error) {
	return p.db.Create(photo).Error
}

// DeletePhotoByID implements photo.PhotoRepository.
func (p *pgRepository) DeletePhotoByID(id uint) (err error) {
	return p.db.Delete(&photo.Photo{}, id).Error
}

// GetAllPhotosByUserID implements photo.PhotoRepository.
func (p *pgRepository) GetAllPhotosByUserID(userId uint) (photos []photo.Photo, err error) {
	return photos, p.db.Where(&photo.Photo{UserID: userId}).Find(&photos).Error
}

// GetPhotoByID implements photo.PhotoRepository.
func (p *pgRepository) GetPhotoByID(id uint) (photo photo.Photo, err error) {
	return photo, p.db.First(&photo, id).Error
}

// UpdatePhoto implements photo.PhotoRepository.
func (p *pgRepository) UpdatePhoto(photo *photo.Photo) (err error) {
	return p.db.Updates(photo).Error
}

// Create implements socialmedia.SocialMediaRepository.
func (p *pgRepository) CreateSocialMedia(socialMedia *socialmedia.SocialMedia) (err error) {
	return p.db.Create(socialMedia).Error
}

// Delete implements socialmedia.SocialMediaRepository.
func (p *pgRepository) DeleteSocialMediaByID(id uint) (err error) {
	return p.db.Delete(&socialmedia.SocialMedia{}, id).Error
}

// GetAll implements socialmedia.SocialMediaRepository.
func (p *pgRepository) GetAllSocialMediasByUserID(userId uint) (socialMedias []socialmedia.SocialMedia, err error) {
	return socialMedias, p.db.Where(&socialmedia.SocialMedia{UserID: userId}).Find(&socialMedias).Error
}

// GetByID implements socialmedia.SocialMediaRepository.
func (p *pgRepository) GetSocialMediaByID(id uint) (socialMedia socialmedia.SocialMedia, err error) {
	return socialMedia, p.db.First(&socialMedia, id).Error
}

// Update implements socialmedia.SocialMediaRepository.
func (p *pgRepository) UpdateSocialMedia(socialMedia *socialmedia.SocialMedia) (err error) {
	return p.db.Updates(socialMedia).Error
}

// GetUserByEmail implements user.UserRepository.
func (p *pgRepository) GetUserByEmail(email string) (user user.User, err error) {
	return user, p.db.Where("email = ?", email).First(&user).Error
}

// CreateUser implements user.UserRepository.
func (p *pgRepository) CreateUser(user *user.User) (err error) {
	return p.db.Create(user).Error
}

func NewUserPostgresRepository(db *gorm.DB) user.UserRepository {
	return &pgRepository{db}
}
func NewSocialMediaPostgresRepository(db *gorm.DB) socialmedia.SocialMediaRepository {
	return &pgRepository{db}
}
func NewPhotoPostgresRepository(db *gorm.DB) photo.PhotoRepository {
	return &pgRepository{db}
}
func NewCommentPostgresRepository(db *gorm.DB) comment.CommentRepository {
	return &pgRepository{db}
}
