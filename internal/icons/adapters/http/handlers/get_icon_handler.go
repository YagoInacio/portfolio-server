package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/icons/usecases"
)

type GetIconHandler struct {
	GetIconUseCase usecases.GetIconUseCase
}

func NewGetIconHandler(getIconUseCase usecases.GetIconUseCase) *GetIconHandler {
	return &GetIconHandler{
		GetIconUseCase: getIconUseCase,
	}
}

func (h *GetIconHandler) Handler(c echo.Context) error {
	iconName := c.Param("name")
	file, err := h.GetIconUseCase.Execute(iconName)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.Blob(200, "image/png", file)
}
