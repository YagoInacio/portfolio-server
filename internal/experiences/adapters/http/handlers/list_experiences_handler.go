package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/experiences/usecases"
)

type ListExperiencesHandler struct {
	ListExperiencesUseCase usecases.ListExperiencesUseCase
}

func NewListExperiencesHandler(listExperiencesUseCase usecases.ListExperiencesUseCase) *ListExperiencesHandler {
	return &ListExperiencesHandler{
		ListExperiencesUseCase: listExperiencesUseCase,
	}
}

// @Tags Experiences
// @Router	/experiences [get]
// @Summary List Experiences
// @Description  This endpoint lists all registered experiences
// @Produce      json
// @Success      200  {array}  usecases.ListExperiencesOutput
func (h *ListExperiencesHandler) Handler(c echo.Context) error {
	result, err := h.ListExperiencesUseCase.Execute()

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
