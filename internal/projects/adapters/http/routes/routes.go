package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/projects/adapters/http/handlers"
	repositories "github.com/yagoinacio/portfolio-server/internal/projects/adapters/persistance/database"
	"github.com/yagoinacio/portfolio-server/internal/projects/usecases"
	tech_repositories "github.com/yagoinacio/portfolio-server/internal/technologies/adapters/persistance/database"
	"github.com/yagoinacio/portfolio-server/pkg/database/mongodb"
)

func SetupRoutes(e *echo.Group) {
	collection := mongodb.GetCollection("projects")
	tech_collection := mongodb.GetCollection("technologies")
	projectsRepository := repositories.NewProjectsRepository(collection)
	techRepository := tech_repositories.NewTechnologiesRepository(tech_collection)

	create := usecases.NewCreateProjectUseCase(projectsRepository, techRepository)
	e.POST("", handlers.NewCreateProjectHandler(*create).Handler)
}
