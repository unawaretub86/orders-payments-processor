package usecase_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"

	"orders-payments-processor/internal/domain/entities"
	"orders-payments-processor/internal/domain/usecase"
	"orders-payments-processor/internal/domain/usecase/mocks"
)

func TestCreateOrder_Success(t *testing.T) {
	mockRepo := &mocks.Mocks{}

	validBody := `{
			"user_id": "1234pruebaCompleta",
			"item": "pruebaCompleta",
			"quantity": 1111,
			"total_price": 2222
		}`
	requestID := "1234567890"

	mockRepo.ConvertOrderRequestFunc = func(order *entities.OrderRequest, requestId string) (*string, error) {
		userID := "user_id"
		return &userID, nil
	}

	err := usecase.ConvertOrderRequest(validBody, requestID)

	mockSQS := mocks.NewMockSQS("us-east-2")
	queueURL := "https://queue.amazonaws.com/80398EXAMPLE/MyQueue"

	messageAttributes := map[string]*sqs.MessageAttributeValue{
		"Source": {
			DataType:    aws.String("String"),
			StringValue: aws.String("order-processor-events"),
		},
	}

	_, err = mockSQS.SendMessage(&sqs.SendMessageInput{
		MessageBody:       aws.String(`{"order_id": "1234567890"}`),
		QueueUrl:          &queueURL,
		MessageAttributes: messageAttributes,
	})
	if err != nil {
		t.Errorf("Error sending message: %v", err)
	}

	if err != nil {
		t.Errorf("Error updating payment: %v", err)
	}

	assert.NoError(t, err)
}

func TestCreatePayment_Error(t *testing.T) {
	mockRepo := &mocks.Mocks{}

	invalidBody := `{invalid_field: 1234567890}`
	requestID := "1234567890"

	expectedError := fmt.Errorf("invalid input data")

	mockRepo.ConvertOrderRequestFunc = func(order *entities.OrderRequest, requestId string) (*string, error) {
		userID := "user_id"
		return &userID, nil
	}

	err := usecase.ConvertOrderRequest(invalidBody, requestID)

	assert.Error(t, err, expectedError)
}
