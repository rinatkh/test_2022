package usecase

import (
	"github.com/rinatkh/test_2022/config"
	"github.com/rinatkh/test_2022/internal/constants"
	"github.com/rinatkh/test_2022/internal/pets"
	"github.com/rinatkh/test_2022/internal/photo"
	"github.com/rinatkh/test_2022/pkg/utils"
	"github.com/sirupsen/logrus"
	"time"
)

type PetsUseCase struct {
	cfg       *config.Config
	log       *logrus.Entry
	repoPets  pets.PetsRepository
	repoPhoto photo.PhotoRepository
}

func NewPetsUC(cfg *config.Config, log *logrus.Entry, repoPets pets.PetsRepository, repoPhoto photo.PhotoRepository) pets.UseCase {
	return &PetsUseCase{
		cfg:       cfg,
		log:       log,
		repoPets:  repoPets,
		repoPhoto: repoPhoto,
	}
}

func (u PetsUseCase) CreatePets(params *pets.CreatePetsParams) (*pets.CreatePetsResponse, error) {
	id, err := utils.GenUUID()
	if err != nil {
		return nil, err
	}
	pet := pets.Pets{
		Id:        id,
		Name:      params.Name,
		Age:       params.Age,
		Type:      params.Type,
		CreatedAt: time.Now(),
	}
	err = u.repoPets.CreatePet(&pet)
	if err != nil {
		return nil, err
	}

	return &pets.CreatePetsResponse{Pets: pet}, nil
}

func (u PetsUseCase) GetPets(params *pets.GetPetsParams) (*pets.GetPetsResponse, error) {
	res, err := u.repoPets.GetPets(params)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return &pets.GetPetsResponse{Count: 0, Items: nil}, nil
	}
	for i, j := range *res {
		ph, err := u.repoPhoto.GetPhotoByPetsId(j.Id)
		if err != nil {
			return nil, err
		}
		if ph == nil {
			continue
		}
		(*res)[i].Photos = *ph
	}
	return &pets.GetPetsResponse{Count: int64(len(*res)), Items: *res}, nil
}

func (u PetsUseCase) DeletePets(params *pets.DeletePetsParams) (*pets.DeletePetsResponse, error) {
	var result pets.DeletePetsResponse
	for _, i := range params.IDs {
		res, err := u.repoPets.GetPetById(i)
		if err != nil {
			return nil, err
		}
		if res == nil {
			result.Errors = append(result.Errors, pets.Error{
				Id:    i,
				Error: constants.ErrIdPetsNotFound,
			})
			continue
		}
		err = u.repoPets.DeletePet(i)
		if err != nil {
			result.Errors = append(result.Errors, pets.Error{
				Id:    i,
				Error: err.Error(),
			})
			continue
		}
		err = u.repoPhoto.DeletePhotoByPetsId(i)
		if err != nil {
			return nil, err
		}
		result.Deleted += 1
	}
	return &result, nil
}
