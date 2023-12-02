package handlers

import (
	"mime"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/images/usecases"
)

type GetImageHandler struct {
	GetImageUseCase usecases.GetImageUseCase
}

func NewGetImageHandler(getImageUseCase usecases.GetImageUseCase) *GetImageHandler {
	return &GetImageHandler{
		GetImageUseCase: getImageUseCase,
	}
}

func (h *GetImageHandler) Handler(c echo.Context) error {
	imageName := c.Param("name")
	file, err := h.GetImageUseCase.Execute(imageName)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	contentType := mime.TypeByExtension("." + imageName[strings.LastIndex(imageName, ".")+1:])

	return c.Blob(200, contentType, file)
}
