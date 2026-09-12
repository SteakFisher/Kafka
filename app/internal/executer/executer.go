package executer

import (
	"github.com/codecrafters-io/kafka-starter-go/app/internal/parser"
	"github.com/codecrafters-io/kafka-starter-go/app/internal/response"
)

func Execute(req parser.Payload) response.ResponsePayload {
	resp := response.NewResponsePayload(req.Header.Correlation_id)

	if !isValidApiVersion(req.Header.Request_api_key) {
		resp.Error(35)
	}

	return *resp
}

func isValidApiVersion(apiVersion int16) bool {
	return apiVersion <= 4 && apiVersion >= 0
}
