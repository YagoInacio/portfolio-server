package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/pkg/icons/adapters/http/handlers"
	"github.com/yagoinacio/portfolio-server/pkg/icons/usecases"
)

func SetupRoutes(e *echo.Group) {

	getIcon := usecases.GetIconUseCase{}
	e.GET("/:name", handlers.NewGetIconHandler(getIcon).Handler)
}
