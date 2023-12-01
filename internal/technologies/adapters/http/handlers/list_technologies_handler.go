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

// @Tags Technologies
// @Router	/technologies [get]
// @Summary List Technologies
// @Description  This endpoint lists all technologies that are marked as display: true
// @Produce      json
// @Success      200  {array}  usecases.ListTechnologiesOutput
func (h *ListTechnologiesHandler) Handler(c echo.Context) error {
	result, err := h.ListTechnologiesUseCase.Execute()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
