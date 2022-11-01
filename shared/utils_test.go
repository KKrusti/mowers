package shared

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToInt(t *testing.T) {
	number := "5"

	got := StringToInt(number)

	assert.Equal(t, 5, got)
}

func TestStringToInt_withError(t *testing.T) {
	number := "?"

	got := StringToInt(number)

	assert.Equal(t, 0, got)
}
