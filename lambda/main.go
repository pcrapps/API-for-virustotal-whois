package main

import (
	"context"
	"fmt"
	"lambdaapi/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func router(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	lookuptype := req.QueryStringParameters["lookuptype"]
	if lookuptype == "whois" {
		fmt.Println("Whois Type")
		return handlers.WhoisHandler(ctx, req)
	}
	if lookuptype == "virustotal" {
		return handlers.VirusTotalHandler(ctx, req)
	}
	return nil, nil
}

func main() {
	lambda.Start(router)
}
