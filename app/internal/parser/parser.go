package parser

import (
	"encoding/binary"

	"github.com/codecrafters-io/kafka-starter-go/app/internal/domain"
)

type RequestHeader struct {
	Request_api_key     domain.ApiRequest
	Request_api_version int16
	Correlation_id      int32
	Client_id           []byte
	Tag_buffer          byte
}

type RequestBody interface {
	Type() domain.ApiRequest
	Parse([]byte) error
}

type Payload struct {
	Message_size uint32
	Header       RequestHeader
	Body         RequestBody
}

func Parse(buf []byte, size uint32) Payload {
	client_id_len := int16(binary.BigEndian.Uint16(buf[8:10]))
	request_api_version := domain.ApiRequest(binary.BigEndian.Uint16(buf[0:2]))
	var body RequestBody

	switch request_api_version {
	case domain.ApiVersions:
		body = &APIVersionRequest{}
	}

	body.Parse(buf[client_id_len+11:])

	return Payload{
		Message_size: size,
		Header: RequestHeader{
			Request_api_key:     request_api_version,
			Request_api_version: int16(binary.BigEndian.Uint16(buf[2:4])),
			Correlation_id:      int32(binary.BigEndian.Uint32(buf[4:8])),
			Client_id:           buf[10 : client_id_len+10],
		},
		Body: body,
	}
}
