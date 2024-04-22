package validation

import (
	"BE-shop/models/dto/json"
	"BE-shop/models/dto/trxItemsDto"
)

func TrxItemsValidation(req trxItemsDto.TrxItemsReq) []json.ValidationField {
	var validationErrors []json.ValidationField

	if req.ItemCode == "" || len(req.ItemCode) > 10 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "code",
			Message:   "Code items cannot be empty and max 10",
		})
	}

	if req.Quantity <= 0 || req.Quantity > 10000 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "amount",
			Message:   "Amount must be greater than zero and max 10000",
		})
	}

	if req.TransactionType != "IN" && req.TransactionType != "OUT" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "transaction_type",
			Message:   "TransactionType must be either IN or OUT",
		})
	}

	return validationErrors
}
