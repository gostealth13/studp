package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//Connect udp
	conn, err := net.Dial("udp", "localhost:543")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//simple write
	conn.Write([]byte("1,50.123,54.334"))

	//simple Read
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received ", string(buffer[0:n]))

}
