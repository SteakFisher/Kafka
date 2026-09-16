package response

import (
	"encoding/binary"

	"github.com/codecrafters-io/kafka-starter-go/app/internal/domain"
)

type ResponseHeader struct {
	correlation_id int32
}

type ResponseBody interface {
	Type() domain.ApiRequest
	Serialize(buf []byte) int
}

type ResponsePayload struct {
	message_size uint32
	header       ResponseHeader
	error_code   int16
	body         ResponseBody
}

func NewResponsePayload(correlation_id int32) *ResponsePayload {
	return &ResponsePayload{
		message_size: 6,
		header: ResponseHeader{
			correlation_id: correlation_id,
		},
	}
}

func (r *ResponsePayload) Error(code int16) {
	r.error_code = code
}

func (r *ResponsePayload) SetApiVersionResponse(api_keys []ApiKeyVersioning) {
	r.body = &ApiVersionResponse{
		api_keys: api_keys,
	}
}

func (r *ResponsePayload) Serialize() []byte {
	buf := make([]byte, 128)

	i := 4
	// binary.BigEndian.PutUint32(buf[0:4], uint32(r.message_size))
	binary.BigEndian.PutUint32(buf[i:i+4], uint32(r.header.correlation_id))

	i += 4
	binary.BigEndian.PutUint16(buf[i:i+2], uint16(r.error_code))

	i += 2

	i += r.body.Serialize(buf[i:])
	binary.BigEndian.PutUint32(buf[0:4], uint32(i-4))

	return buf[:i]
}

// func (r ResponsePayload) SetHeaders(header ResponseHeader, message_size uint32) {
// 	r.message_size = message_size
// 	r.header = header
// }
