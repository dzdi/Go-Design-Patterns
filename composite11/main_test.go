package composite11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOrganization(t *testing.T) {
	count := NewOrganization().Count()
	assert.Equal(t, 20, count)
}
