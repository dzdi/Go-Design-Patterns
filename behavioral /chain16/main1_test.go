package chain16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSensitiveWordFilterChain_AddFilter(t *testing.T) {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	assert.Equal(t, false, chain.Filter("test"))

	chain.AddFilter(&PoliticalWordFilter{})
	assert.Equal(t, true, chain.Filter("test"))

}
