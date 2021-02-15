package main

import (
	"context"
	"lambdaapi/handlers"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func router(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	lookuptype := req.QueryStringParameters["lookuptype"]

	var bodystring string

	if lookuptype == "vt" {
		hash := req.QueryStringParameters["hash"]
		if hash != "" {
			s, _ := handlers.VirusTotalURLHandler(hash)
			bodystring = s
		}
		domain := req.QueryStringParameters["domain"]
		s, _ := handlers.VirusTotalURLHandler(domain)
		bodystring = s

	}
	if lookuptype == "whois" {
		domain := req.QueryStringParameters["domain"]
		s, _ := handlers.WhoisHandler(domain)
		bodystring = s
	}
	resp := events.APIGatewayProxyResponse{}
	resp.Body = string(bodystring)
	resp.StatusCode = http.StatusOK
	resp.IsBase64Encoded = false
	resp.Headers = map[string]string{"Content-Type": "application/json"}
	return &resp, nil
}

func main() {
	lambda.Start(router)
}
