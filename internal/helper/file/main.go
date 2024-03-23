package helper

import (
	"fmt"
	"os"
)

var osStat = os.Stat

func RemoveFile(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println(err)
	}
}

func CheckIfPathExists(path string) (bool, error) {
	_, err := osStat(path)
	
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
