package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/experiences/adapters/http/handlers"
	repositories "github.com/yagoinacio/portfolio-server/internal/experiences/adapters/persistance/database"
	"github.com/yagoinacio/portfolio-server/internal/experiences/usecases"
	tech_repositories "github.com/yagoinacio/portfolio-server/internal/technologies/adapters/persistance/database"
	"github.com/yagoinacio/portfolio-server/pkg/database/mongodb"
)

func SetupRoutes(e *echo.Group) {
	collection := mongodb.GetCollection("experiences")
	tech_collection := mongodb.GetCollection("technologies")
	xpRepository := repositories.NewExperiencesRepository(collection)
	techRepository := tech_repositories.NewTechnologiesRepository(tech_collection)

	create := usecases.NewCreateExperienceUseCase(xpRepository, techRepository)
	e.POST("", handlers.NewCreateExperienceHandler(*create).Handler)
}
