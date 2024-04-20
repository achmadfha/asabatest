package validation

import (
	"BE-shop/models/dto/itemsDto"
	"BE-shop/models/dto/json"
)

func ValidationItems(req itemsDto.ItemsRequest) []json.ValidationField {
	var validationErrors []json.ValidationField

	if req.Code == "" || len(req.Code) > 10 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "code",
			Message:   "Code items cannot be empty and max 10",
		})
	}

	if req.Name == "" || len(req.Name) > 50 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "name",
			Message:   "Name cannot be empty and max 50",
		})
	}

	if req.Amount <= 0 || req.Amount > 10000 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "amount",
			Message:   "Amount must be greater than zero and max 10000",
		})
	}

	if req.Description == "" || len(req.Description) > 150 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "description",
			Message:   "Description cannot be empty and max 150",
		})
	}

	if req.StatusActive != true && req.StatusActive != false {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "status_active",
			Message:   "StatusActive must be either true or false",
		})
	}

	return validationErrors
}
