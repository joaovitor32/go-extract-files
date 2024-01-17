package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfCompacted(t *testing.T) {

	isZip := CheckIfCompacted("aquivo.zip")

	assert.True(t, isZip, "Check if file is zip, --FAILED")

}

func TestUnzip(t *testing.T) {

}
