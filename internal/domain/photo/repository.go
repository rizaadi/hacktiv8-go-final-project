package photo

type PhotoRepository interface {
	CreatePhoto(photo *Photo) (err error)
	GetPhotoByID(id uint) (photo Photo, err error)
	GetAllPhotosByUserID(userId uint) (photos []Photo, err error)
	UpdatePhoto(photo *Photo) (err error)
	DeletePhotoByID(id uint) (err error)
}
