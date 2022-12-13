package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/internal/middleware"
)

func MapPetsRoutes(router fiber.Router, h *PetsHandler, mw *middleware.MDWManager) {

	router.Post("/pets", mw.VerifyMiddleware(), h.CreatePets())
	router.Get("/pets", mw.VerifyMiddleware(), h.GetPets())
	router.Delete("/pets", mw.VerifyMiddleware(), h.DeletePets())
}
