package photo

type service struct {
	repo PhotoRepository
}

// CreatePhoto implements Service.
func (s *service) CreatePhoto(photo *Photo) (err error) {
	return s.repo.CreatePhoto(photo)
}

// DeletePhotoByID implements Service.
func (s *service) DeletePhotoByID(id uint) (err error) {
	return s.repo.DeletePhotoByID(id)
}

// GetAllPhotosByUserID implements Service.
func (s *service) GetAllPhotosByUserID(userId uint) (photos []Photo, err error) {
	return s.repo.GetAllPhotosByUserID(userId)
}

// GetPhotoByID implements Service.
func (s *service) GetPhotoByID(id uint) (photo Photo, err error) {
	return s.repo.GetPhotoByID(id)
}

// UpdatePhoto implements Service.
func (s *service) UpdatePhoto(photo *Photo) error {
	if err := s.repo.UpdatePhoto(photo); err != nil {
		return err
	}
	updatedPhoto, err := s.repo.GetPhotoByID(photo.ID)
	if err != nil {
		return err
	}
	*photo = updatedPhoto
	return nil
}

type Service interface {
	CreatePhoto(photo *Photo) (err error)
	GetPhotoByID(id uint) (photo Photo, err error)
	GetAllPhotosByUserID(userId uint) (photos []Photo, err error)
	UpdatePhoto(photo *Photo) error
	DeletePhotoByID(id uint) (err error)
}

func NewService(repo PhotoRepository) Service {
	return &service{repo: repo}

}
