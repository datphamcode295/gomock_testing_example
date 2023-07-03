package main_test

import (
	"errors"
	"fmt"
	"testing"

	"example.com/testing"
	mock_main "example.com/testing/mocks"
	"github.com/golang/mock/gomock"
)

func TestCharge(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockPaymentProcessor := mock_main.NewMockPaymentProcessor(mockCtrl)
	testPaymentProcessorClient := &main.PaymentProcessorClient{PaymentProcessor: mockPaymentProcessor}

	defer mockCtrl.Finish()

	mockPaymentProcessor.EXPECT().Charge(100.0, "test_token").Return(nil).Times(1)

	err := testPaymentProcessorClient.Charge(100.0, "test_token")

	if err != nil {
		t.Fail()
	}

	err = testPaymentProcessorClient.Charge(10.0, "test_token")

	if err.Error() != "Charge too low" {
		fmt.Println(fmt.Errorf("Error: %w", err))
		t.Fail()
	}
}

func TestChargeTooLow(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockPaymentProcessor := mock_main.NewMockPaymentProcessor(mockCtrl)
	testPaymentProcessorClient := &main.PaymentProcessorClient{PaymentProcessor: mockPaymentProcessor}

	defer mockCtrl.Finish()

	mockPaymentProcessor.EXPECT().Charge(gomock.Any(), "test_token").Return(errors.New("stripe")).AnyTimes()

	err := testPaymentProcessorClient.Charge(10.0, "test_token")

	if err.Error() != "Charge too low" {
		fmt.Println(fmt.Errorf("Error: %w", err))
		t.Fail()
	}

	err = testPaymentProcessorClient.Charge(20.0, "test_token")

	if err.Error() != "stripe" {
		fmt.Println(fmt.Errorf("Error: %w", err))
		t.Fail()
	}
}
