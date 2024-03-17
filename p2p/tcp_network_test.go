package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	// Test Start method of TCPNetwork
	tn := &TCPNetwork{}
	err := tn.Start()
	defer tn.Close()

	assert.NoError(t, err)
	assert.NotNil(t, tn.listener)
}
