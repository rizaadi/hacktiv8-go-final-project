package comment

type service struct {
	repo CommentRepository
}

// CreateComment implements Service.
func (s *service) CreateComment(comment *Comment) (err error) {
	return s.repo.CreateComment(comment)
}

// DeleteCommentByID implements Service.
func (s *service) DeleteCommentByID(id uint) (err error) {
	return s.repo.DeleteCommentByID(id)
}

// GetAllCommentsByPhotoID implements Service.
func (s *service) GetAllCommentsByPhotoID(photoID uint) (comments []Comment, err error) {
	return s.repo.GetAllCommentsByPhotoID(photoID)
}

// GetCommentByID implements Service.
func (s *service) GetCommentByID(id uint) (comment Comment, err error) {
	return s.repo.GetCommentByID(id)
}

// UpdateComment implements Service.
func (s *service) UpdateComment(comment *Comment) (err error) {
	if err := s.repo.UpdateComment(comment); err != nil {
		return err
	}
	updateComment, err := s.repo.GetCommentByID(comment.ID)
	if err != nil {
		return err
	}
	*comment = updateComment
	return nil
}

type Service interface {
	CreateComment(comment *Comment) (err error)
	GetCommentByID(id uint) (comment Comment, err error)
	GetAllCommentsByPhotoID(photoID uint) (comments []Comment, err error)
	UpdateComment(comment *Comment) (err error)
	DeleteCommentByID(id uint) (err error)
}

func NewService(repo CommentRepository) Service {
	return &service{repo: repo}

}
