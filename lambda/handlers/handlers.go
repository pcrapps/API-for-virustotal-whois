package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"
)

func WhoisHandler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	domain := req.QueryStringParameters["domain"]

	result, _ := whois.Whois(domain)
	parsResult, _ := whoisparser.Parse(result)

	resp := events.APIGatewayProxyResponse{}
	stringBody, _ := json.Marshal(parsResult)
	resp.Body = string(stringBody)
	resp.StatusCode = http.StatusOK
	resp.IsBase64Encoded = false
	resp.Headers = map[string]string{"Content-Type": "application/json"}
	return &resp, nil
}
