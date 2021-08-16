package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_quantity_less_than_one(t *testing.T) {
	// Arrange
	request := NewTransactionRequest{
		Quantity: 0,
	}

	// Act
	AppError := request.Validate()

	// Assert
	if AppError.Message != "Product quantity must more than 0" {
		t.Error("Invalid message when testing transaction quantity validation")
	}

	if AppError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code when testing transaction quantity validation")
	}
}
