package xutil

import (
	"encoding/json"
	"fmt"
)

type PaymentCancellationStatus int

const (
	Ongoing = iota
	Cancelled
	CancellationError
)

type PaymentCancellationResponse struct {
	Status PaymentCancellationStatus
}

// Serialization functions

func (p PaymentCancellationResponse) Marshal() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Failed to marshal Cancellation Response: %s", err)
	}
	return bytes
}

func UnmarshalPaymentCancellationResponse(bytes []byte) (PaymentCancellationResponse, error) {
	res := PaymentCancellationResponse{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		fmt.Printf("Failed to unmarshal Response: %s", err)
		return res, err
	}
	return res, nil
}
