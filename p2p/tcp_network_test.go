package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPNetwork_Start(t *testing.T) {
	// Create a new TCPNetwork instance
	tn := NewTCPNetwork()

	// Start the TCP network
	err := tn.Start()

	// Assert that there are no errors
	assert.NoError(t, err)

	// Assert that the listener is not nil
	assert.NotNil(t, tn.Listener)
}
