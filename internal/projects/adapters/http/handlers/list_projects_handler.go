package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/projects/usecases"
)

type ListProjectsHandler struct {
	ListProjectsUseCase usecases.ListProjectsUseCase
}

func NewListProjectsHandler(listProjectsUseCase usecases.ListProjectsUseCase) *ListProjectsHandler {
	return &ListProjectsHandler{
		ListProjectsUseCase: listProjectsUseCase,
	}
}

// @Tags Projects
// @Router	/projects [get]
// @Summary List projects
// @Description  This endpoint lists all projects that are marked as display: true
// @Produce      json
// @Success      200  {array}  usecases.ListProjectsOutput
func (h *ListProjectsHandler) Handler(c echo.Context) error {
	result, err := h.ListProjectsUseCase.Execute()

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
