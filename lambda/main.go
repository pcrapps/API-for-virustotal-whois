package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"
)

//MyEvent event intake
type MyEvent struct {
	Domain string `json:"domain"`
}

func x(ctx context.Context, input MyEvent) (whoisparser.WhoisInfo, error) {
	result, err := whois.Whois(input.Domain)
	if err != nil {
		log.Fatal(err)
	}
	parsResult, err := whoisparser.Parse(result)
	if err != nil {
		log.Fatal(err)
	}

	return parsResult, nil

}

func main() {

	lambda.Start(x)
}
