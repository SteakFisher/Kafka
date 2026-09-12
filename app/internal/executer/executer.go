package executer

import (
	"github.com/codecrafters-io/kafka-starter-go/app/internal/parser"
	"github.com/codecrafters-io/kafka-starter-go/app/internal/response"
)

func Execute(req parser.Payload) response.ResponsePayload {
	resp := response.NewResponsePayload(req.Header.Correlation_id)

	return resp
}
