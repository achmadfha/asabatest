package itemsDto

import (
	"github.com/google/uuid"
	"time"
)

type (
	Items struct {
		ItemsID         uuid.UUID `json:"items_id"`
		Code            string    `json:"code"`
		Name            string    `json:"name"`
		Amount          int       `json:"amount"`
		Description     string    `json:"description"`
		StatusActive    bool      `json:"status_active"`
		TransactionType string    `json:"transaction_type"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	}

	ItemsRequest struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Amount       int    `json:"amount"`
		Description  string `json:"description"`
		StatusActive bool   `json:"status_active"`
	}

	ItemsResponse struct {
		ItemsID      uuid.UUID `json:"items_id"`
		Code         string    `json:"code"`
		Name         string    `json:"name"`
		Amount       int       `json:"amount"`
		Description  string    `json:"description"`
		StatusActive bool      `json:"status_active"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
