package helper

import (
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

func TestUnzip(t *testing.T) {
	var path = "dummy-path"
	var destiny = "dummy-destiny"

	_, err := Unzip(path, destiny)

	if assert.NotNil(t, err) {
		assert.Equal(t, "unarr: open file failed", err.Error())
	}
}
