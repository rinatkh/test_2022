package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/internal/photo"
	"github.com/sirupsen/logrus"
)

type PhotoHandler struct {
	photoUC photo.UseCase
	log     *logrus.Entry
}

func NewPhotoHandler(photoUC photo.UseCase, log *logrus.Entry) *PhotoHandler {

	return &PhotoHandler{
		photoUC: photoUC,
		log:     log,
	}
}

func (u PhotoHandler) UploadPhoto() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		image, err := ctx.FormFile("image")
		if err != nil {
			return err
		}
		data, err := u.photoUC.UploadPhoto(image, ctx.Params("id"))
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
