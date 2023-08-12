package main

import (
	"orders-payments-processor/internal/domain/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.HttpHandler)
}
