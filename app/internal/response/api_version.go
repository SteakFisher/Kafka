package response

import (
	"encoding/binary"

	"github.com/codecrafters-io/kafka-starter-go/app/internal/domain"
)

type ApiKeyVersioning struct {
	Api_key     int16
	Min_version int16
	Max_version int16
	Tag_buffer  byte
}

type ApiVersionResponse struct {
	api_keys         []ApiKeyVersioning
	throttle_time_ms int32
	tag_buffer       byte
}

func (ApiVersionResponse) Type() domain.ApiRequest {
	return domain.ApiVersions
}

func (r *ApiVersionResponse) Serialize(buf []byte) int {
	api_keys_len := len(r.api_keys) + 1
	buf[0] = byte(api_keys_len)
	i := 1
	for j := range api_keys_len - 1 {
		i += r.api_keys[j].Serialize(buf[i:])
	}
	binary.BigEndian.PutUint32(buf[i:i+4], uint32(r.throttle_time_ms))
	i += 4
	buf[i] = r.tag_buffer
	i += 1
	return i
}

func (r *ApiKeyVersioning) Serialize(buf []byte) int {
	binary.BigEndian.PutUint16(buf[0:2], uint16(r.Api_key))
	binary.BigEndian.PutUint16(buf[2:4], uint16(r.Min_version))
	binary.BigEndian.PutUint16(buf[4:6], uint16(r.Max_version))
	buf[6] = r.Tag_buffer
	return 7
}

var _ ResponseBody = (*ApiVersionResponse)(nil)
