package usecases

import (
	"os"

	file_storage "github.com/yagoinacio/portfolio-server/pkg/file_storage/firebase"
)

type GetImageUseCase struct{}

// @Tags Images
// @Router	/images/{name} [get]
// @Summary Get image file
// @Description  This endpoint returns the image file by its name
// @Param name   path      string  true  "File name"
// @Produce      image/png
// @Produce      image/jpeg
// @Produce      image/svg+xml
// @Produce      image/gif
// @Success      200  {file} file
func (g *GetImageUseCase) Execute(input string) ([]byte, error) {
	filePath := "images/" + input
	result, err := file_storage.DownloadFileIntoMemory(os.Stdout, "my-first-project-9872e.appspot.com", filePath)

	if err != nil {
		return nil, err
	}

	return result, nil
}
