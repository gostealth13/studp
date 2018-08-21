package server

import (
	"log"
	"net"
)

type Config struct {
}

type callback func(pc net.PacketConn)

// Struct of player
type sv struct {
	c *Config
}

// Server interface is used for defining all methods
type Server interface {
	Run(c callback)
}

// Contructor
func New(c *Config) Server {
	return &sv{c}
}

func (s *sv) Run(c callback) {
	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", "0.0.0.0:543")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()
	for {
		c(pc)
	}
}
