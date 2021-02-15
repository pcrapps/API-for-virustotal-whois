package handlers

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"

	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"

	"github.com/williballenthin/govt"
)

//WhoisHandler is a wrapper for th whois website via a golang package
//it takes the aws gateway request and returns a response
func WhoisHandler(domain string) (string, error) {

	result, _ := whois.Whois(domain)
	parsResult, _ := whoisparser.Parse(result)
	stringBody, _ := json.Marshal(parsResult)

	return string(stringBody), nil
}

//VirusTotalURLHandler is a wrapper for th virus total website via a golang package
//it takes the aws gateway request and returns a response
func VirusTotalURLHandler(domain string) (string, error) {
	apikey := getAWSSecret()
	apiurl := "https://www.virustotal.com/vtapi/v2/"

	client, _ := govt.New(govt.SetApikey(apikey), govt.SetUrl(apiurl))
	domainReport, _ := client.GetDomainReport(domain)
	jsonReport, _ := json.Marshal(domainReport)
	return string(jsonReport), nil
}

//VirusTotalHashHandler is a wrapper for th virus total website via a golang package
//it takes the aws gateway request and returns a response
func VirusTotalHashHandler(hash string) (string, error) {
	apikey := getAWSSecret()
	apiurl := "https://www.virustotal.com/vtapi/v2/"

	client, _ := govt.New(govt.SetApikey(apikey), govt.SetUrl(apiurl))
	fileReport, _ := client.GetFileReport(hash)
	jsonReport, _ := json.Marshal(fileReport)
	return string(jsonReport), nil
}

func getAWSSecret() string {
	secretName := "virustotalapi"
	region := "us-east-2"

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}
	result, _ := svc.GetSecretValue(input)
	return *result.SecretString

}
