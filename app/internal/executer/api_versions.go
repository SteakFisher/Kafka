package executer

import (
	"github.com/codecrafters-io/kafka-starter-go/app/internal/parser"
	"github.com/codecrafters-io/kafka-starter-go/app/internal/response"
)

func api_versions(req parser.Payload, resp *response.ResponsePayload) *response.ResponsePayload {
	resp.SetApiVersionResponse([]response.ApiKeyVersioning{
		{
			Api_key:     18,
			Min_version: 0,
			Max_version: 4,
		},
		{
			Api_key:     75,
			Min_version: 0,
			Max_version: 0,
		},
	})

	return resp
}
