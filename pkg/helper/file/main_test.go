package helper

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfPathExists(t *testing.T) {
	var path = "dummy-path"

	// Save current function and restore at the end:
	oldOsStat := osStat
	defer func() { osStat = oldOsStat }()

	myOsStat := func(name string) (fs.FileInfo, error) {
		return nil, nil
	}

	osStat = myOsStat

	pathExist, _ := CheckIfPathExists(path)
	assert.True(t, pathExist, "Check if path exists, --FAILED")
}
