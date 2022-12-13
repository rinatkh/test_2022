package usecase

import (
	"github.com/rinatkh/test_2022/config"
	"github.com/rinatkh/test_2022/internal/constants"
	"github.com/rinatkh/test_2022/internal/pets"
	"github.com/rinatkh/test_2022/internal/photo"
	"github.com/rinatkh/test_2022/pkg/utils"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type PhotoUseCase struct {
	cfg       *config.Config
	log       *logrus.Entry
	repoPets  pets.PetsRepository
	repoPhoto photo.PhotoRepository
}

func NewPhotoUC(cfg *config.Config, log *logrus.Entry, repoPets pets.PetsRepository, repoPhoto photo.PhotoRepository) photo.UseCase {
	return &PhotoUseCase{
		cfg:       cfg,
		log:       log,
		repoPets:  repoPets,
		repoPhoto: repoPhoto,
	}
}

func (u PhotoUseCase) UploadPhoto(fileHeader *multipart.FileHeader, id string) (*photo.UploadPhotoResponse, error) {
	res, err := u.repoPets.GetPetById(id)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, constants.ErrDBNotFound
	}
	src, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	uuid, err := utils.GenUUID()
	if err != nil {
		return nil, err
	}

	filename := uuid + filepath.Ext(fileHeader.Filename)

	dst, err := os.Create("/opt/pics/" + filename)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	url := "/opt/pics/" + filename

	ph := photo.Photo{
		Id:        uuid,
		Url:       url,
		PetsId:    id,
		CreatedAt: time.Now(),
	}
	err = u.repoPhoto.CreatePhoto(&ph)
	if err != nil {
		return nil, err
	}

	return &photo.UploadPhotoResponse{Photo: photo.Photo{
		Id:  id,
		Url: url,
	}}, nil
}
