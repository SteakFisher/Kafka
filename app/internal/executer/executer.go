package executer

import (
	"github.com/codecrafters-io/kafka-starter-go/app/internal/domain"
	"github.com/codecrafters-io/kafka-starter-go/app/internal/parser"
	"github.com/codecrafters-io/kafka-starter-go/app/internal/response"
)

func Execute(req parser.Payload) response.ResponsePayload {
	resp := response.NewResponsePayload(req.Header.Correlation_id)

	if !isValidApiVersion(req.Header.Request_api_version) {
		resp.Error(35)
	}

	switch req.Header.Request_api_key {
	case domain.ApiVersions:
		api_versions(req, resp)
	}

	return *resp
}

func isValidApiVersion(apiVersion int16) bool {
	return apiVersion <= 4 && apiVersion >= 0
}
