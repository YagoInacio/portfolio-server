package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	icon_routes "github.com/yagoinacio/portfolio-server/pkg/icons/adapters/http/routes"
	tech_routes "github.com/yagoinacio/portfolio-server/pkg/technologies/adapters/http/routes"
)

func NewServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, duration=${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	api := e.Group("/api")

	technologies := api.Group("/technologies")
	tech_routes.SetupRoutes(technologies)

	icons := api.Group("/icons")
	icon_routes.SetupRoutes(icons)

	return e
}
