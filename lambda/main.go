package main

import (
	"lambdaapi/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(handlers.WhoisHandler)
}
