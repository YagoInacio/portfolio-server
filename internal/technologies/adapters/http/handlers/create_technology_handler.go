package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/technologies/usecases"
)

type CreateTechnologyHandler struct {
	CreateTechnologyUseCase usecases.CreateTechnologyUseCase
}

func NewCreateTechnologyHandler(createTechnologyUseCase usecases.CreateTechnologyUseCase) *CreateTechnologyHandler {
	return &CreateTechnologyHandler{
		CreateTechnologyUseCase: createTechnologyUseCase,
	}
}

// @Tags Technologies
// @Router	/technologies [post]
// @Summary creates a new technology
// @Description  This endpoint creates a new technology
// @Accept       json
// @Produce      json
// @Param request body usecases.TechnologyInput true "creation attributes"
func (h *CreateTechnologyHandler) Handler(c echo.Context) error {
	request := new(usecases.TechnologyInput)
	if err := c.Bind(request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := h.CreateTechnologyUseCase.Execute(*request)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
