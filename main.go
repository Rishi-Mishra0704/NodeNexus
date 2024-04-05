package main

import (
	"fmt"

	"github.com/Rishi-Mishra0704/NodeNexus/p2p"
)

func main() {
	s := p2p.NewServer()
	s.Commands = make(chan p2p.Command)
	fmt.Println("Starting server on :8080")
}
