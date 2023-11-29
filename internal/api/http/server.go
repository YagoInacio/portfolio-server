package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yagoinacio/portfolio-server/pkg/icons/handlers"
	"github.com/yagoinacio/portfolio-server/pkg/icons/usecases"
)

func NewServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	uc := usecases.GetIconUseCase{}

	handler := handlers.NewGetIconHandler(uc)

	// add routes
	e.GET("/icons/:name", handler.Handler)

	return e
}
