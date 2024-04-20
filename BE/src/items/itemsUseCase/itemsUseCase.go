package itemsUseCase

import (
	"BE-shop/models/dto/itemsDto"
	"BE-shop/src/items"
	"github.com/google/uuid"
	"time"
)

type itemsUC struct {
	itemsRepository items.ItemsRepository
}

func NewItemsUseCase(items items.ItemsRepository) items.ItemsUseCase {
	return &itemsUC{items}
}

func (i itemsUC) CreateItem(items itemsDto.ItemsRequest) (itemsDto.ItemsResponse, error) {
	itemsID, err := uuid.NewRandom()
	if err != nil {
		return itemsDto.ItemsResponse{}, err
	}

	trxType := "IN"
	currentTime := time.Now()

	itemsData := itemsDto.Items{
		ItemsID:         itemsID,
		Code:            items.Code,
		Name:            items.Name,
		Amount:          items.Amount,
		Description:     items.Description,
		StatusActive:    items.StatusActive,
		TransactionType: trxType,
		CreatedAt:       currentTime,
		UpdatedAt:       currentTime,
	}

	err = i.itemsRepository.CreateItem(itemsData)
	if err != nil {
		return itemsDto.ItemsResponse{}, err
	}

	itemsResponse := itemsDto.ItemsResponse{
		ItemsID:      itemsID,
		Code:         items.Code,
		Name:         items.Name,
		Amount:       items.Amount,
		Description:  items.Description,
		StatusActive: items.StatusActive,
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
	}

	return itemsResponse, nil
}

func (i itemsUC) RetrieveAllItems(transactionType string) ([]itemsDto.Items, error) {
	itemsData, err := i.itemsRepository.RetrieveAllItems(transactionType)
	if err != nil {
		return []itemsDto.Items{}, err
	}

	return itemsData, nil
}
