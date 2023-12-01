package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/technologies/usecases"
)

type AlterTechnologyDisplayHandler struct {
	AlterTechnologyDisplayUseCase usecases.AlterTechnologyDisplayUseCase
}

func NewAlterTechnologyDisplayHandler(alterTechnologyDisplayUseCase usecases.AlterTechnologyDisplayUseCase) *AlterTechnologyDisplayHandler {
	return &AlterTechnologyDisplayHandler{
		AlterTechnologyDisplayUseCase: alterTechnologyDisplayUseCase,
	}
}

func (h *AlterTechnologyDisplayHandler) Handler(c echo.Context) error {
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

	result, err := h.AlterTechnologyDisplayUseCase.Execute(usecases.AlterTechnologyDisplayInput{
		ID:      inputId,
		Display: inputDisplayAsBoolean,
	})
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
