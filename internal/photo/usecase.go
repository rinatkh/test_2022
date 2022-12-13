package photo

import "mime/multipart"

type UseCase interface {
	UploadPhoto(fileHeader *multipart.FileHeader, id string) (*UploadPhotoResponse, error)
}
