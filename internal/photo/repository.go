package photo

type PhotoRepository interface {
	GetPhotoById(id string) (*Photo, error)
	DeletePhoto(id string) error
	CreatePhoto(photo *Photo) error
	GetPhotoByPetsId(id string) (*[]Photo, error)
	DeletePhotoByPetsId(id string) error
}
