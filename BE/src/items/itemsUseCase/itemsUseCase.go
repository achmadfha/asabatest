package itemsUseCase

import (
	"BE-shop/models/dto/itemsDto"
	"BE-shop/src/items"
	"errors"
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

func (i itemsUC) RetrieveAllItems() ([]itemsDto.ItemsResponse, error) {
	itemsData, err := i.itemsRepository.RetrieveAllItems()
	if err != nil {
		return []itemsDto.ItemsResponse{}, err
	}

	return itemsData, nil
}

func (i itemsUC) RetrieveItemsByCode(code string) (itemsDto.Items, error) {
	itemsData, err := i.itemsRepository.RetrieveItemsByCode(code)
	if err != nil {
		if err.Error() == "01" {
			return itemsDto.Items{}, errors.New("01")
		}
		return itemsDto.Items{}, err
	}

	return itemsData, nil
}

func (i itemsUC) UpdateItemsByCode(items itemsDto.ItemsUpdate) (itemsDto.Items, error) {
	itemsData, err := i.itemsRepository.RetrieveItemsByCode(items.Code)
	if err != nil {
		if err.Error() == "01" {
			return itemsDto.Items{}, errors.New("01")
		}
		return itemsDto.Items{}, err
	}

	oldAmount := itemsData.Amount

	if items.Amount != 0 {
		itemsData.Amount = items.Amount
	}

	if items.StatusActive {
		itemsData.StatusActive = items.StatusActive
	}

	if items.TransactionType != "" {
		itemsData.TransactionType = items.TransactionType
	}

	if itemsData.TransactionType == "IN" {
		itemsData.Amount += oldAmount
	} else if itemsData.TransactionType == "OUT" {
		if itemsData.Amount > oldAmount {
			return itemsDto.Items{}, errors.New("02")
		}
		itemsData.Amount = oldAmount - itemsData.Amount
	}

	itemsUpdate := itemsDto.ItemsUpdate{
		Code:            itemsData.Code,
		Amount:          itemsData.Amount,
		StatusActive:    itemsData.StatusActive,
		TransactionType: itemsData.TransactionType,
	}

	err = i.itemsRepository.UpdateItemsByCode(itemsUpdate)
	if err != nil {
		return itemsDto.Items{}, err
	}

	return itemsData, nil
}

func (i itemsUC) DeleteItemsByCode(code string) error {
	itemsData, err := i.itemsRepository.RetrieveItemsByCode(code)
	if err != nil {
		if err.Error() == "01" {
			return errors.New("01")
		}
		return err
	}

	err = i.itemsRepository.DeleteItemsByCode(itemsData.Code)
	if err != nil {
		return err
	}

	return nil
}
