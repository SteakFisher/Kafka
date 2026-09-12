package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Just the first 12 bytes of the request:
	// length = 38
	// API key = 18
	// API version = 4
	// correlation ID = 7
	req := make([]byte, 12)

	binary.BigEndian.PutUint32(req[0:4], 38)
	binary.BigEndian.PutUint16(req[4:6], 18)
	binary.BigEndian.PutUint16(req[6:8], 4)
	binary.BigEndian.PutUint32(req[8:12], 7)

	fmt.Printf("sending: % x\n", req)

	_, err = conn.Write(req)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("received %d bytes: % x\n", n, buf[:n])
}
