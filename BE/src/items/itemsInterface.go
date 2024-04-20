package items

import "BE-shop/models/dto/itemsDto"

type ItemsRepository interface {
	CreateItem(items itemsDto.Items) error
	RetrieveAllItems(transactionType string) ([]itemsDto.Items, error)
	RetrieveItemsByCode(code string) (itemsDto.Items, error)
	UpdateItemsByCode(items itemsDto.ItemsUpdate) error
	DeleteItemsByCode(code string) error
}

type ItemsUseCase interface {
	CreateItem(items itemsDto.ItemsRequest) (itemsDto.ItemsResponse, error)
	RetrieveAllItems(transactionType string) ([]itemsDto.Items, error)
	RetrieveItemsByCode(code string) (itemsDto.Items, error)
	UpdateItemsByCode(items itemsDto.ItemsUpdate) (itemsDto.Items, error)
	DeleteItemsByCode(code string) error
}
