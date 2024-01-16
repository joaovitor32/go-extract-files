package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func validate(item string) bool {
	return item != "c"
}

func TestFilter(t *testing.T) {
	array := []string{"a", "b", "c"}
	result := Filter(array, validate)

	expected := []string{"a", "b"}

	assert.Equal(t, expected, result, "Check if array was filtered, --FAILED")

}
