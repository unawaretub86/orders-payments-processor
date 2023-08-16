package mocks

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"

	"orders-payments-processor/internal/domain/entities"
)

type Mocks struct {
	ConvertOrderRequestFunc func(order *entities.OrderRequest, requestId string) (*string, error)
}

type MockSQS struct {
	sqsiface.SQSAPI
	messages map[string][]*sqs.Message
	region   string
}

func NewMockSQS(region string) *MockSQS {
	return &MockSQS{
		messages: make(map[string][]*sqs.Message),
		region:   "us-east-2",
	}
}

func (m *MockSQS) SendMessage(in *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	if _, ok := m.messages[*in.QueueUrl]; !ok {
		m.messages[*in.QueueUrl] = []*sqs.Message{}
	}

	m.messages[*in.QueueUrl] = append(m.messages[*in.QueueUrl], &sqs.Message{
		Body: in.MessageBody,
	})
	return &sqs.SendMessageOutput{}, nil
}

func (m *Mocks) ConvertOrderRequest(order *entities.OrderRequest, requestId string) (*string, error) {
	return m.ConvertOrderRequestFunc(order, requestId)
}
