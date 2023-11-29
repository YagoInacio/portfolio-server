package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/pkg/icons/usecases"
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
	path, err := h.GetIconUseCase.Execute(iconName)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.File(path)
}
