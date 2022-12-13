package httpServer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/internal/middleware"
	"github.com/rinatkh/test_2022/pkg/storage"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	serverLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	petsHTTP "github.com/rinatkh/test_2022/internal/pets/delivery/http"
	petsRepository "github.com/rinatkh/test_2022/internal/pets/repository"
	petsUsecase "github.com/rinatkh/test_2022/internal/pets/usecase"

	photoHTTP "github.com/rinatkh/test_2022/internal/photo/delivery/http"
	photoRepository "github.com/rinatkh/test_2022/internal/photo/repository"
	photoUsecase "github.com/rinatkh/test_2022/internal/photo/usecase"
)

func (s *Server) MapHandlers(app *fiber.App) error {

	postgreConnection, err := storage.InitPsqlDB(s.cfg)
	if err != nil {
		return err
	}

	petsRepo := petsRepository.NewPostgresRepository(postgreConnection, s.log)
	photoRepo := photoRepository.NewPostgresRepository(postgreConnection, s.log)

	petsUC := petsUsecase.NewPetsUC(s.cfg, s.log, petsRepo, photoRepo)
	photoUC := photoUsecase.NewPhotoUC(s.cfg, s.log, petsRepo, photoRepo)

	petsHandler := petsHTTP.NewPetsHandler(petsUC, s.log)
	photoHandler := photoHTTP.NewPhotoHandler(photoUC, s.log)

	app.Use(serverLogger.New())
	if _, ok := os.LookupEnv("LOCAL"); !ok {
		app.Use(recover.New())
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	mw := middleware.NewMDWManager(s.cfg, petsUC)

	petsHTTP.MapPetsRoutes(app, petsHandler, mw)
	photoHTTP.MapPhotoRoutes(app, photoHandler, mw)

	return nil
}
