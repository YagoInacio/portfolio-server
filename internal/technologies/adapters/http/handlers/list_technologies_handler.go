package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/technologies/usecases"
)

type ListTechnologiesHandler struct {
	ListTechnologiesUseCase usecases.ListTechnologiesUseCase
}

func NewListTechnologiesHandler(listTechnologiesUseCase usecases.ListTechnologiesUseCase) *ListTechnologiesHandler {
	return &ListTechnologiesHandler{
		ListTechnologiesUseCase: listTechnologiesUseCase,
	}
}

func (h *ListTechnologiesHandler) Handler(c echo.Context) error {
	result, err := h.ListTechnologiesUseCase.Execute()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
