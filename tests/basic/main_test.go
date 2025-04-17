package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	// got := AddOne(1)
	// want := 2
	// if got != want {
	// 	t.Errorf("AddOne(%d) = %d; want %d", 1, got, want)
	// }

	assert.Equal(t, 3, AddOne(1), "AddOne(1) should be 2")
}
