package trxItemsDto

import (
	"github.com/google/uuid"
	"time"
)

type (
	TrxItems struct {
		TrxID     uuid.UUID `json:"transaction_id"`
		ItemCode  string    `json:"items_code"`
		TrxType   string    `json:"transaction_type"`
		Amount    int       `json:"amount"`
		Quantity  int       `json:"quantity"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	TrxItemsReq struct {
		ItemCode        string `json:"items_code"`
		TransactionType string `json:"transaction_type"`
		Quantity        int    `json:"quantity"`
	}

	TrxItemsRes struct {
		TrxID     uuid.UUID `json:"transaction_id"`
		ItemCode  string    `json:"items_code"`
		TrxType   string    `json:"transaction_type"`
		Quantity  int       `json:"quantity"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	TrxUpdateReq struct {
		TrxID    uuid.UUID `json:"transaction_id"`
		ItemCode string    `json:"items_code"`
		TrxType  string    `json:"transaction_type"`
		Quantity int       `json:"quantity"`
		Amount   int       `json:"amount"`
	}

	TrxItemsUpdateReq struct {
		TrxID           uuid.UUID `json:"transaction_id"`
		ItemCode        string    `json:"items_code"`
		TransactionType string    `json:"transaction_type"`
		Quantity        int       `json:"quantity"`
	}
)
