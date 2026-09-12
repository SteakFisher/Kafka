package parser

import "encoding/binary"

type RequestHeader struct {
	Request_api_key     int16
	Request_api_version int16
	Correlation_id      int32
	Client_id           []byte
}

type Payload struct {
	Message_size uint32
	Header       RequestHeader
}

func Parse(buf []byte, size uint32) Payload {
	return Payload{
		Message_size: size,
		Header: RequestHeader{
			Request_api_key:     int16(binary.BigEndian.Uint16(buf[0:2])),
			Request_api_version: int16(binary.BigEndian.Uint16(buf[2:4])),
			Correlation_id:      int32(binary.BigEndian.Uint32(buf[4:8])),
		},
	}
}
