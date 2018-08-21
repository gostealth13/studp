package main

import (
	"fmt"
	"log"
	"net"

	"github.com/wiltongarcia/studp/server"
)

func main() {
	c := &server.Config{}
	s := server.New(c)
	s.Run(func(pc net.PacketConn) {

		buffer := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buffer)
		if err != nil {
			log.Fatal("")
		}
		fmt.Println("Received ", string(buffer[0:n]), " from ", addr)

		pc.WriteTo([]byte("Hello from server"), addr)
	})
}
