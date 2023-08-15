package usecase

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"orders-payments-processor/internal/domain/entities"
)

func ConvertOrderRequest(body, requestId string) error {
	var orderRequest entities.OrderRequest
	err := json.Unmarshal([]byte(body), &orderRequest)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error marshaling API Gateway request: %v]", err, requestId)
		return err
	}

	return sendSQS(orderRequest, requestId)
}

func sendSQS(order entities.OrderRequest, requestId string) error {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	sqsClient := sqs.New(sess)

	queueURL := os.Getenv("SQS_URL")

	orderJSON, err := json.Marshal(order)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error marshaling order request: %v]", err, requestId)
		return err
	}

	messageAttributes := map[string]*sqs.MessageAttributeValue{
		"Source": {
			DataType:    aws.String("String"),
			StringValue: aws.String("orders-processor"),
		},
	}

	_, err = sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody:       aws.String(string(orderJSON)),
		QueueUrl:          &queueURL,
		MessageAttributes: messageAttributes,
	})

	if err != nil {
		fmt.Printf("[RequestId: %s][Error sending message to SQS: %v]", err, requestId)
		return err
	}

	return nil
}
