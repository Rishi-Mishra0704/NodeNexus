package p2p

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() *TCPNetwork {
	tn := NewTCPNetwork()
	err := tn.Start()
	if err != nil {
		log.Fatalf("Failed to start TCP network: %v", err)
	}
	return tn
}

func teardown(tn *TCPNetwork) {
	err := tn.Close()
	if err != nil {
		log.Fatalf("Failed to close TCP network: %v", err)
	}
}

func TestTCPNetwork_Start(t *testing.T) {
	tn := setup()
	defer teardown(tn)

	assert.NotNil(t, tn.Listener)
}

func TestTCPNetwork_Close(t *testing.T) {
	tn := setup()
	defer teardown(tn)

	err := tn.Close()
	assert.NoError(t, err)
}
