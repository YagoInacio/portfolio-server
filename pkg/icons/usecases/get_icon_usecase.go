package usecases

import (
	"os"
	"path/filepath"
	"strings"
)

type GetIconUseCase struct{}

func (g *GetIconUseCase) Execute(input string) (string, error) {
	var result string
	err := filepath.Walk("./icons", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Verificar se o nome do arquivo corresponde e não possui extensão
		if strings.HasPrefix(info.Name(), input+".") && !info.IsDir() {
			result = path
			return filepath.SkipDir // Pule o diretório, pois já encontramos o arquivo
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return result, nil
}
