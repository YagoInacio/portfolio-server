package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/technologies/adapters/http/handlers"
	repositories "github.com/yagoinacio/portfolio-server/internal/technologies/adapters/persistance/database"
	usecases "github.com/yagoinacio/portfolio-server/internal/technologies/usecases"
	"github.com/yagoinacio/portfolio-server/pkg/database/mongodb"
)

func SetupRoutes(e *echo.Group) {
	collection := mongodb.GetCollection("technologies")
	techRepository := repositories.NewTechnologiesRepository(collection)

	create := usecases.NewCreateTechnologyUseCase(techRepository)
	e.POST("", handlers.NewCreateTechnologyHandler(*create).Handler)

	list := usecases.NewListTechnologiesUseCase(techRepository)
	e.GET("", handlers.NewListTechnologiesHandler(*list).Handler)
}
