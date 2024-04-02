package p2p

import (
	"encoding/gob"
	"io"
)

// Decoder is an interface for decoding messages.
type Decoder interface {
	Decode(io.Reader, *Message) error
}

// GobDecoder decodes messages using gob encoding.
type GobDecoder struct{}

// Decode decodes a message using gob encoding.
func (gd GobDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

// DefaultDecoder decodes messages using a default decoding mechanism.
type DefaultDecoder struct{}

// Decode decodes a message using a default decoding mechanism.
func (dd DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1028) // Buffer to read the message payload
	n, err := r.Read(buf)     // Read from the reader
	if err != nil {
		return err
	}
	msg.Payload = buf[:n]
	return nil
}
