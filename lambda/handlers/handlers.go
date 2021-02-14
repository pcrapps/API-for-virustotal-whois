package handlers

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"

	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"

	vt "github.com/VirusTotal/vt-go"
)

//WhoisHandler is a wrapper for th whois website via a golang package
//it takes the aws gateway request and returns a response
func WhoisHandler(domain string) (string, error) {

	result, _ := whois.Whois(domain)
	parsResult, _ := whoisparser.Parse(result)
	stringBody, _ := json.Marshal(parsResult)

	return string(stringBody), nil
}

//VirusTotalHandler is a wrapper for th whois website via a golang package
//it takes the aws gateway request and returns a response
func VirusTotalHandler(filehash string) (string, error) {

	apikey := getAWSSecret
	client := vt.NewClient(apikey())
	file, err := client.GetObject(vt.URL("files/%s", filehash))
	if err != nil {
		log.Fatal(err)
	}
	stringBody, _ := json.Marshal(file)
	return string(stringBody), nil

}

func getAWSSecret() string {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("virustotalapi"),
	}
	result, _ := svc.GetSecretValue(input)
	return *result.SecretString

}
