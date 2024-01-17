package helper

import (
	"fmt"
	helper "go-extract-files/pkg/helper/file"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Close() (err error) {
	return nil
}

func TestCheckIfCompacted(t *testing.T) {

	isZip := CheckIfCompacted("aquivo.zip")

	assert.True(t, isZip, "Check if file is zip, --FAILED")

}

func TestUnzipFails(t *testing.T) {
	var path = "dummy-path"
	var destiny = "dummy-destiny"

	_, err := Unzip(path, destiny)

	if assert.NotNil(t, err) {
		assert.Equal(t, "unarr: open file failed", err.Error())
	}
}

func TestUnzipSuccess(t *testing.T) {
	path, _ := filepath.Abs("../../../testdata/teste.zip")
	var destiny = "test-output"

	expectedResponse := []string{"DWSample.zip"}

	response, _ := Unzip(path, destiny)

	fmt.Print("aqui", response)

	assert.Equal(t, response, expectedResponse, "Check if internalzip array has content, --FAILED")
	helper.RemoveFile(destiny)
}
