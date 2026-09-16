package handler

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/app/internal/executer"
	"github.com/codecrafters-io/kafka-starter-go/app/internal/parser"
)

func Handle(conn net.Conn) {
	for {
		size_buf := make([]byte, 4)

		_, err := conn.Read(size_buf)

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

		request_payload := parser.Parse(buf, size)
		response_payload := executer.Execute(request_payload)

		conn.Write(response_payload.Serialize())
	}
}
