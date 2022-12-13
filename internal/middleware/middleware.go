package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/config"
	"github.com/rinatkh/test_2022/internal/constants"
	"github.com/rinatkh/test_2022/internal/pets"
)

type MDWManager struct {
	cfg     *config.Config
	usersUC pets.UseCase
}

func NewMDWManager(cfg *config.Config, usersUC pets.UseCase) *MDWManager {
	return &MDWManager{
		cfg:     cfg,
		usersUC: usersUC,
	}
}

func (mw *MDWManager) VerifyMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		key := c.Get("X-API-KEY")
		if key != mw.cfg.Keys.ApiKey {
			return constants.AuthError
		}

		return c.Next()
	}
}
