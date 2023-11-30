package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/database/mongodb"
	"github.com/yagoinacio/portfolio-server/pkg/technologies/adapters/http/handlers"
	repositories "github.com/yagoinacio/portfolio-server/pkg/technologies/adapters/persistance/database"
	usecases "github.com/yagoinacio/portfolio-server/pkg/technologies/usecases"
)

func SetupRoutes(e *echo.Group) {
	collection := mongodb.GetCollection("technologies")
	techRepository := repositories.NewTechnologiesRepository(collection)

	create := usecases.NewCreateTechnologyUseCase(techRepository)
	e.POST("", handlers.NewCreateTechnologyHandler(*create).Handler)
}
