package handler

import (
	"context"
	"fmt"

	"orders-payments-processor/internal/domain/usecase"

	"github.com/aws/aws-lambda-go/events"
)

func HttpHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	requestID := fmt.Sprintf("%v", ctx.Value("aws_request_id"))

	fmt.Println("ctx", ctx)
	fmt.Println("request", request)

	body := request.Body
	fmt.Println("body", body)

	response := events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       "{\"message\": \"Order Created \"}",
	}

	return response, usecase.CreateOrderRequest(body, requestID)
}
