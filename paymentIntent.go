package xutil

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type PaymentIntent struct {
	To       string
	Uuid     string
	Currency string
	Amount   *big.Float
}

// Amount conversion helpers

func (p *PaymentIntent) AmountFloat64() float64 {
	v, _ := p.Amount.Float64()
	return v
}

func (p *PaymentIntent) AmountInt64() int64 {
	v, _ := p.Amount.Int64()
	return v
}

func (p *PaymentIntent) AmountBigInt() *big.Int {
	vInt64, _ := p.Amount.Int64()
	vBig := big.NewInt(vInt64)
	return vBig
}

// Serialization functions

func (p PaymentIntent) Marshal() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Failed to marshal payment intent: %s", err)
	}
	return bytes
}

func UnmarshalPaymentIntent(bytes []byte) (PaymentIntent, error) {
	res := PaymentIntent{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		fmt.Printf("Failed to unmarshal response: %s", err)
		return res, err
	}
	return res, nil
}
