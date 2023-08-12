package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"orders-payments-processor/internal/domain/handler"
)

func main() {
	lambda.Start(handler.HttpHandler)
}
