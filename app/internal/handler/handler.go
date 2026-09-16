package handler

import (
	"net"

	"github.com/codecrafters-io/kafka-starter-go/app/internal/executer"
	"github.com/codecrafters-io/kafka-starter-go/app/internal/parser"
)

func Handle(conn net.Conn, buf []byte, size uint32) {
	request_payload := parser.Parse(buf, size)
	response_payload := executer.Execute(request_payload)

	conn.Write(response_payload.Serialize())
	conn.Close()
}
