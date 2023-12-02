package handlers

import (
	"mime"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/portfolio-server/internal/images/usecases"
)

type GetIconHandler struct {
	GetIconUseCase usecases.GetIconUseCase
}

func NewGetIconHandler(getIconUseCase usecases.GetIconUseCase) *GetIconHandler {
	return &GetIconHandler{
		GetIconUseCase: getIconUseCase,
	}
}

// @Tags Images
// @Router	/images/icons/{name} [get]
// @Summary Get icon file
// @Description  This endpoint returns the icon file by its name
// @Param name   path      string  true  "File name"
// @Produce      image/png
// @Produce      image/jpeg
// @Produce      image/svg+xml
// @Produce      image/gif
// @Success      200  {file} file
func (h *GetIconHandler) Handler(c echo.Context) error {
	iconName := c.Param("name")
	file, err := h.GetIconUseCase.Execute(iconName)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	contentType := mime.TypeByExtension("." + iconName[strings.LastIndex(iconName, ".")+1:])

	return c.Blob(200, contentType, file)
}
