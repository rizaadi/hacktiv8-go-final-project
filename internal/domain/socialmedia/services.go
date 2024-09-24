package socialmedia

type service struct {
	repo SocialMediaRepository
}

// CreateSocialMedia implements Service.
func (s *service) CreateSocialMedia(socialMedia *SocialMedia) (err error) {
	return s.repo.CreateSocialMedia(socialMedia)
}

// DeleteSocialMediaByID implements Service.
func (s *service) DeleteSocialMediaByID(id uint) (err error) {
	return s.repo.DeleteSocialMediaByID(id)
}

// GetAllSocialMediasByUserID implements Service.
func (s *service) GetAllSocialMediasByUserID(userId uint) (socialMedias []SocialMedia, err error) {
	return s.repo.GetAllSocialMediasByUserID(userId)
}

// GetSocialMediaByID implements Service.
func (s *service) GetSocialMediaByID(id uint) (socialMedia SocialMedia, err error) {
	return s.repo.GetSocialMediaByID(id)
}

// UpdateSocialMedia implements Service.
func (s *service) UpdateSocialMedia(socialMedia *SocialMedia) error {
	if err := s.repo.UpdateSocialMedia(socialMedia); err != nil {
		return err
	}
	updatedSocialMedia, err := s.repo.GetSocialMediaByID(socialMedia.ID)
	if err != nil {
		return err
	}
	*socialMedia = updatedSocialMedia
	return nil
}

type Service interface {
	CreateSocialMedia(socialMedia *SocialMedia) (err error)
	GetSocialMediaByID(id uint) (socialMedia SocialMedia, err error)
	GetAllSocialMediasByUserID(userId uint) (socialMedias []SocialMedia, err error)
	UpdateSocialMedia(socialMedia *SocialMedia) error
	DeleteSocialMediaByID(id uint) (err error)
}

func NewService(repo SocialMediaRepository) Service {
	return &service{repo: repo}
}
