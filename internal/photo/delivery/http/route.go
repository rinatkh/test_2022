package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/internal/middleware"
)

func MapPhotoRoutes(router fiber.Router, h *PhotoHandler, mw *middleware.MDWManager) {
	router.Post("/pets/:id/photo", mw.VerifyMiddleware(), h.UploadPhoto())
}
