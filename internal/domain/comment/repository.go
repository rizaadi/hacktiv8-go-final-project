package comment

type CommentRepository interface {
	CreateComment(comment *Comment) (err error)
	GetCommentByID(id uint) (comment Comment, err error)
	GetAllCommentsByPhotoID(photoID uint) (comments []Comment, err error)
	UpdateComment(comment *Comment) (err error)
	DeleteCommentByID(id uint) (err error)
}
