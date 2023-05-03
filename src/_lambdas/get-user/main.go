package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"test-go-lambda/handlers"
)

func main() {
	lambda.Start(handlers.GetUser)
}
