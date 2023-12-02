package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/experiences/usecases"
)

type CreateExperienceHandler struct {
	CreateExperienceUseCase usecases.CreateExperienceUseCase
}

func NewCreateExperienceHandler(createExperienceUseCase usecases.CreateExperienceUseCase) *CreateExperienceHandler {
	return &CreateExperienceHandler{
		CreateExperienceUseCase: createExperienceUseCase,
	}
}

// @Tags Experiences
// @Router	/experiences [post]
// @Summary Creates an experience
// @Description  This endpoint creates an experience
// @Accept       json
// @Produce      json
// @Param request body usecases.CreateExperienceInput true "creation attributes"
// @Success      201  {object}  usecases.CreateExperienceOutput
func (h *CreateExperienceHandler) Handler(c echo.Context) error {
	request := new(usecases.CreateExperienceInput)
	if err := c.Bind(request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := h.CreateExperienceUseCase.Execute(*request)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
