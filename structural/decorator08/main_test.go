package decorator08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorSquare_Draw(t *testing.T) {
	s := Square{}
	square := NewColorSquare(s, "red")
	draw := square.Draw()
	assert.Equal(t, "this is a Square, color is red", draw)
}
