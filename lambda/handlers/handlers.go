package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"

	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"

	vt "github.com/VirusTotal/vt-go"
)

//WhoisHandler is a wrapper for th whois website via a golang package
//it takes the aws gateway request and returns a response
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

//VirusTotalHandler is a wrapper for th whois website via a golang package
//it takes the aws gateway request and returns a response
func VirusTotalHandler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	apikey := getAWSSecret
	filehash := req.QueryStringParameters["filehash"]

	client := vt.NewClient(apikey())

	file, err := client.GetObject(vt.URL("files/%s", filehash))
	if err != nil {
		log.Fatal(err)
	}

	resp := events.APIGatewayProxyResponse{}
	stringBody, _ := json.Marshal(file)
	resp.Body = string(stringBody)
	resp.StatusCode = http.StatusOK
	resp.IsBase64Encoded = false
	resp.Headers = map[string]string{"Content-Type": "application/json"}
	return &resp, nil

}

func getAWSSecret() string {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("virustotalapi"),
	}
	result, _ := svc.GetSecretValue(input)
	return *result.SecretString

}
