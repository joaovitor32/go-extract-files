package helper

import (
	"fmt"
	helper "go-extract-files/pkg/helper/file"
	utils "go-extract-files/pkg/utils/array"
	"path/filepath"

	unarr "github.com/gen2brain/go-unarr"
)

var newArchive = unarr.NewArchive

func CheckIfCompacted(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".zip" || ext == ".rar" || ext == ".7"
}

func Unzip(path string, destiny string) (contents []string, err error) {
	r, err := newArchive(filepath.Join(path))

	if err != nil {
		return nil, err
	}

	defer r.Close()

	extractedFiles, err := r.Extract(destiny)

	if err != nil {
		fmt.Print(err)
		return []string{}, err
	}

	return utils.Filter(extractedFiles, CheckIfCompacted), nil
}

func UnzipAll(path string, destiny string) {
	internalZips, err := Unzip(path, destiny)

	if err != nil {
		panic(err)
	}

	if len(internalZips) == 0 {
		fmt.Print("\n- Não há nenhum zip interno para o arquivo -> ", path)
		return
	}

	for _, internalFile := range internalZips {
		UnzipAll(destiny+"/"+internalFile, destiny)
		helper.RemoveFile(destiny + "/" + internalFile)
	}
}
