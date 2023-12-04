package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/projects/usecases"
)

type UpdateProjectHandler struct {
	UpdateProjectUseCase usecases.UpdateProjectUseCase
}

func NewUpdateProjectHandler(updateProjectUseCase usecases.UpdateProjectUseCase) *UpdateProjectHandler {
	return &UpdateProjectHandler{
		UpdateProjectUseCase: updateProjectUseCase,
	}
}

// @Tags Projects
// @Router	/projects [patch]
// @Summary Updates project
// @Description  This endpoint can alter any project field
// @Accept       json
// @Produce      json
// @Param request body usecases.UpdateProjectInput true "update attributes"
// @Success      200  {object}  usecases.UpdateProjectOutput
func (h *UpdateProjectHandler) Handler(c echo.Context) error {
	request := new(usecases.UpdateProjectInput)
	if err := c.Bind(request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := h.UpdateProjectUseCase.Execute(*request)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
