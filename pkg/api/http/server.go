package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	icon_routes "github.com/yagoinacio/portfolio-server/internal/icons/adapters/http/routes"
	tech_routes "github.com/yagoinacio/portfolio-server/internal/technologies/adapters/http/routes"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/yagoinacio/portfolio-server/docs"
)

// @title Portfolio server
// @version 1.0
// @description This API is designed to improve the management of project and experience information for your portfolio website, ensuring a seamless rendering and effortless updating process. Additionally, it facilitates interaction with Firebase Storage, allowing you to efficiently load and display icons and images, enhancing the overall visual experience of your portfolio.

// @contact.name Yago Faran
// @contact.email yagofaran@gmail.com

// @license.name MIT

// @BasePath /api
func NewServer() *echo.Echo {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

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
