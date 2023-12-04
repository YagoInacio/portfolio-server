package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/projects/usecases"
)

type AlterProjectDisplayHandler struct {
	AlterProjectDisplayUseCase usecases.AlterProjectDisplayUseCase
}

func NewAlterProjectDisplayHandler(alterProjectDisplayUseCase usecases.AlterProjectDisplayUseCase) *AlterProjectDisplayHandler {
	return &AlterProjectDisplayHandler{
		AlterProjectDisplayUseCase: alterProjectDisplayUseCase,
	}
}

// @Tags Projects
// @Router	/projects/display/{id} [patch]
// @Summary Alters project display
// @Description  This endpoint alters a project's display parameter
// @Produce      json
// @Param id   path      string  true  "Project Id"
// @Param display query string true "display value to be modified to" Enums(true, false)
// @Success      200  {object}  usecases.AlterProjectDisplayOutput
func (h *AlterProjectDisplayHandler) Handler(c echo.Context) error {
	inputId := c.Param("id")
	inputDisplay := c.QueryParam("display")

	var inputDisplayAsBoolean bool
	if inputDisplay == "true" {
		inputDisplayAsBoolean = true
	} else if inputDisplay == "false" {
		inputDisplayAsBoolean = false
	} else {
		return c.String(http.StatusBadRequest, "invalid display value")
	}

	result, err := h.AlterProjectDisplayUseCase.Execute(usecases.AlterProjectDisplayInput{
		ID:      inputId,
		Display: inputDisplayAsBoolean,
	})
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
