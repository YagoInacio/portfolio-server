package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/experiences/usecases"
)

type UpdateExperienceHandler struct {
	UpdateExperienceUseCase usecases.UpdateExperienceUseCase
}

func NewUpdateExperienceHandler(updateExperienceUseCase usecases.UpdateExperienceUseCase) *UpdateExperienceHandler {
	return &UpdateExperienceHandler{
		UpdateExperienceUseCase: updateExperienceUseCase,
	}
}

// @Tags Experiences
// @Router	/experiences [patch]
// @Summary Updates experience
// @Description  This endpoint can alter any experience field
// @Accept       json
// @Produce      json
// @Param request body usecases.UpdateExperienceInput true "update attributes"
// @Success      200  {object}  usecases.UpdateExperienceOutput
func (h *UpdateExperienceHandler) Handler(c echo.Context) error {
	request := new(usecases.UpdateExperienceInput)
	if err := c.Bind(request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := h.UpdateExperienceUseCase.Execute(*request)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
