package usecase

import (
	"encoding/json"
	"fmt"

	"orders-payments-processor/internal/domain/entities"
)

func CreateOrderRequest(body, requestId string) error {

	fmt.Println(body)

	var orderRequest entities.CreateOrderRequest
	err := json.Unmarshal([]byte(body), &orderRequest)
	if err != nil {
		fmt.Println("Error unmarshaling API Gateway request:", err)
		return err
	}

	fmt.Println(orderRequest, requestId, "hola estoy en la lambda")

	return nil
}
