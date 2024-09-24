package socialmedia

type SocialMediaRepository interface {
	CreateSocialMedia(socialMedia *SocialMedia) (err error)
	GetSocialMediaByID(id uint) (socialMedia SocialMedia, err error)
	GetAllSocialMediasByUserID(userID uint) (socialMedias []SocialMedia, err error)
	UpdateSocialMedia(socialMedia *SocialMedia) (err error)
	DeleteSocialMediaByID(id uint) (err error)
}
