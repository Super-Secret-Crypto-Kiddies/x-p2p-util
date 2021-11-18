package xutil

import (
	"encoding/json"
	"fmt"
)

type PaymentStatus int

const (
	Empty = iota
	Pending
	Verified
	Error
)

type PaymentResponse struct {
	To     string
	Status PaymentStatus
}

// Serialization functions

func (p PaymentResponse) Marshal() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Failed to marshal Response: %s", err)
	}
	return bytes
}

func UnmarshalPaymentResponse(bytes []byte) (PaymentResponse, error) {
	res := PaymentResponse{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		fmt.Printf("Failed to unmarshal Response: %s", err)
		return res, err
	}
	return res, nil
}
