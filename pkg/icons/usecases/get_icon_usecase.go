package usecases

import (
	"os"

	file_storage "github.com/yagoinacio/portfolio-server/internal/file_storage/firebase"
)

type GetIconUseCase struct{}

func (g *GetIconUseCase) Execute(input string) ([]byte, error) {
	filePath := "icons/" + input
	result, err := file_storage.DownloadFileIntoMemory(os.Stdout, "my-first-project-9872e.appspot.com", filePath)

	if err != nil {
		return nil, err
	}

	return result, nil
}
