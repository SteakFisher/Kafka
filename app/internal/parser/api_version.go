package parser

import (
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/app/internal/domain"
)

var _ RequestBody = (*APIVersionRequest)(nil)

type APIVersionRequest struct {
	Client_id        []byte
	Software_version []byte
	Tag_buffer       byte
}

func (APIVersionRequest) Type() domain.ApiRequest {
	return domain.ApiVersions
}

func (r *APIVersionRequest) Parse(buf []byte) error {
	fmt.Printf("%x\n", buf)
	i := 0

	client_id_length := int(buf[i])

	client_id := buf[i : i+client_id_length]
	i += client_id_length

	client_software_version_length := int(buf[i])

	client_software_version := buf[i : i+client_software_version_length] //
	i += client_software_version_length

	tag_buffer := buf[i]

	r.Client_id = client_id
	r.Software_version = client_software_version
	r.Tag_buffer = tag_buffer

	return nil
}
