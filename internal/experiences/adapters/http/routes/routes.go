package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/experiences/adapters/http/handlers"
	repositories "github.com/yagoinacio/portfolio-server/internal/experiences/adapters/persistance/database"
	"github.com/yagoinacio/portfolio-server/internal/experiences/usecases"
	"github.com/yagoinacio/portfolio-server/pkg/database/mongodb"
)

func SetupRoutes(e *echo.Group) {
	collection := mongodb.GetCollection("experiences")
	xpRepository := repositories.NewExperiencesRepository(collection)

	create := usecases.NewCreateExperienceUseCase(xpRepository)
	e.POST("", handlers.NewCreateExperienceHandler(*create).Handler)
}
