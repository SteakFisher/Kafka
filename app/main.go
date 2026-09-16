package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/kafka-starter-go/app/internal/handler"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// TODO: Uncomment the code below to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		size_buf := make([]byte, 4)

		_, err = conn.Read(size_buf)

		if err != nil {
			fmt.Println("Error reading size: ", err)
			conn.Close()
		}

		size := binary.BigEndian.Uint32(size_buf)

		buf := make([]byte, size)
		_, err = conn.Read(buf)

		if err != nil {
			fmt.Println("Error reading size: ", err)

		}
		go handler.Handle(conn, buf, size)
	}
}
