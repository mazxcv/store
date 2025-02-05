package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	listenAddress := "127.0.0.1:4000"
	tr := NewTCPTransport(listenAddress)
	assert.Equal(t, tr.listenAddres, listenAddress)

	// Server
	assert.Nil(t, tr.ListenAndAccept())

}
