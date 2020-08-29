package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hamologist/dice-roll/pkg/app"
)

func main() {
	lambda.Start(app.HandleLambdaRequest)
}