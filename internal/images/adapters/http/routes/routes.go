package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/images/adapters/http/handlers"
	"github.com/yagoinacio/portfolio-server/internal/images/usecases"
)

func SetupRoutes(e *echo.Group) {

	getIcon := usecases.GetIconUseCase{}
	e.GET("/icons/:name", handlers.NewGetIconHandler(getIcon).Handler)

	getImage := usecases.GetImageUseCase{}
	e.GET("/:name", handlers.NewGetImageHandler(getImage).Handler)
}
