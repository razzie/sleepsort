package sleepsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSleepSort(t *testing.T) {
	x := []uint{3, 2, 1}
	SleepSort(x)
	assert.Equal(t, []uint{1, 2, 3}, x)
}
