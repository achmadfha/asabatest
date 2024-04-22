package trxItemsUseCase

import (
	"BE-shop/models/dto/trxItemsDto"
	"BE-shop/src/items"
	"BE-shop/src/trxItems"
	"errors"
	"github.com/google/uuid"
	"time"
)

type trxItemsUC struct {
	trxItemsRepo trxItems.TrxItemsRepository
	items        items.ItemsRepository
}

func NewTrxItemsUseCase(trxItems trxItems.TrxItemsRepository, items items.ItemsRepository) trxItems.TrxItemsUseCase {
	return &trxItemsUC{trxItems, items}
}

func (t trxItemsUC) CreateTrxItems(trx trxItemsDto.TrxItemsReq) (trxItems trxItemsDto.TrxItemsRes, err error) {
	trxID, err := uuid.NewRandom()
	if err != nil {
		return trxItemsDto.TrxItemsRes{}, err
	}
	currentTime := time.Now()

	itemsData, err := t.items.RetrieveItemsByCode(trx.ItemCode)
	if err != nil {
		if err.Error() == "01" {
			return trxItemsDto.TrxItemsRes{}, errors.New("01")
		}
		return trxItemsDto.TrxItemsRes{}, err
	}

	amount := itemsData.Amount
	if trx.TransactionType == "IN" {
		amount += trx.Quantity
	} else if trx.TransactionType == "OUT" {
		if trx.Quantity > amount {
			return trxItemsDto.TrxItemsRes{}, errors.New("02")
		}
		amount -= trx.Quantity
	}

	trxItemsData := trxItemsDto.TrxItems{
		TrxID:     trxID,
		ItemCode:  trx.ItemCode,
		TrxType:   trx.TransactionType,
		Amount:    amount,
		Quantity:  trx.Quantity,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	err = t.trxItemsRepo.CreateTrxItems(trxItemsData)
	if err != nil {
		return trxItemsDto.TrxItemsRes{}, err
	}

	trxItems = trxItemsDto.TrxItemsRes{
		TrxID:     trxID,
		ItemCode:  trx.ItemCode,
		TrxType:   trx.TransactionType,
		Quantity:  trx.Quantity,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	return trxItems, nil
}

func (t trxItemsUC) RetrieveAllTrxItems(transactionType string) (trxItems []trxItemsDto.TrxItemsRes, err error) {
	trxItems, err = t.trxItemsRepo.RetrieveAllTrxItems(transactionType)
	if err != nil {
		return nil, err
	}

	return trxItems, nil
}

func (t trxItemsUC) DeleteTrxItems(trxID string) error {
	err := t.trxItemsRepo.DeleteTrxItems(trxID)
	if err != nil {
		if err.Error() == "01" {
			return errors.New("01")
		}
		return err
	}

	return nil
}
