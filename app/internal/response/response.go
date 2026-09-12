package response

import (
	"encoding/binary"
	"fmt"
)

type ResponseHeader struct {
	correlation_id int32
}

type ResponsePayload struct {
	message_size uint32
	header       ResponseHeader
	error        bool
	error_code   int16
}

func NewResponsePayload(correlation_id int32) *ResponsePayload {
	return &ResponsePayload{
		message_size: 4,
		header: ResponseHeader{
			correlation_id: correlation_id,
		},
	}
}

func (r *ResponsePayload) Error(code int16) {
	r.error = true
	r.error_code = code
	r.message_size += 2
}

func (r *ResponsePayload) Serialize() []byte {
	buf := make([]byte, r.message_size+4)

	binary.BigEndian.PutUint32(buf[0:4], uint32(r.message_size))
	binary.BigEndian.PutUint32(buf[4:8], uint32(r.header.correlation_id))

	if r.error {
		binary.BigEndian.PutUint16(buf[8:10], uint16(r.error_code))
	}

	fmt.Println("Responding with: ", buf)

	return buf
}

// func (r ResponsePayload) SetHeaders(header ResponseHeader, message_size uint32) {
// 	r.message_size = message_size
// 	r.header = header
// }
