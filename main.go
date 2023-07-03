package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type PaymentProcessor interface {
	Charge(amount float64, token string) error
}

type StripePaymentProcessor struct{}

func (s *StripePaymentProcessor) Charge(amount float64, token string) error {
	// charge the credit card
	data := map[string]string{}
	json_data, err := json.Marshal(data)

	resp, err := http.Post("https://api.stripe.com/v1/charges", "", bytes.NewBuffer(json_data))

	if err != nil {
		return err
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return nil
}

type PaymentProcessorClient struct {
	PaymentProcessor PaymentProcessor
}

func (c *PaymentProcessorClient) Charge(amount float64, token string) error {
	if amount < 20 {
		return errors.New("Charge too low")
	}
	return c.PaymentProcessor.Charge(amount, token)
}
