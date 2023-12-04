package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/projects/usecases"
)

type CreateProjectHandler struct {
	CreateProjectUseCase usecases.CreateProjectUseCase
}

func NewCreateProjectHandler(createProjectUseCase usecases.CreateProjectUseCase) *CreateProjectHandler {
	return &CreateProjectHandler{
		CreateProjectUseCase: createProjectUseCase,
	}
}

// @Tags Projects
// @Router	/projects [post]
// @Summary Creates a project
// @Description  This endpoint creates a project
// @Accept       json
// @Produce      json
// @Param request body usecases.CreateProjectInput true "creation attributes"
// @Success      201  {object}  usecases.CreateProjectOutput
func (h *CreateProjectHandler) Handler(c echo.Context) error {
	request := new(usecases.CreateProjectInput)
	if err := c.Bind(request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := h.CreateProjectUseCase.Execute(*request)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
