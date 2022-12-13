package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/internal/constants"
	"github.com/rinatkh/test_2022/internal/pets"
	"github.com/rinatkh/test_2022/pkg/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type PetsHandler struct {
	petsUC pets.UseCase
	log    *logrus.Entry
}

func NewPetsHandler(petsUC pets.UseCase, log *logrus.Entry) *PetsHandler {

	return &PetsHandler{
		petsUC: petsUC,
		log:    log,
	}
}

func (u PetsHandler) CreatePets() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params pets.CreatePetsParams
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.petsUC.CreatePets(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u PetsHandler) DeletePets() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params pets.DeletePetsParams
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.petsUC.DeletePets(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u PetsHandler) GetPets() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params pets.GetPetsParams

		var err error
		params.Limit, err = strconv.ParseInt(ctx.Query("limit", "20"),
			10, 64)
		if err != nil {
			return err
		}
		params.Offset, err = strconv.ParseInt(ctx.Query("offset", "0"),
			10, 64)
		if err != nil {
			return err
		}
		hasPhotos, err := strconv.ParseBool(ctx.Query("has_photos"))
		if err == nil {
			params.HasPhotos = &hasPhotos
		}

		data, err := u.petsUC.GetPets(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
