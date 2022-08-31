package proxy06

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserProxy_Login(t *testing.T) {
	u := &User{}
	proxy := NewUserProxy(u)
	err := proxy.Login("test", "password")

	require.Nil(t, err)
}
