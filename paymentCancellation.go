package xutil

import (
	"encoding/json"
	"fmt"
)

type PaymentCancellation struct {
	Address string
}

func (p PaymentCancellation) Marshal() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Failed to marshal payment cancellation: %s", err)
	}
	return bytes
}

func UnmarshalPaymentCancellation(bytes []byte) (PaymentCancellation, error) {
	res := PaymentCancellation{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		fmt.Printf("Failed to unmarshal response: %s", err)
		return res, err
	}
	return res, nil
}
