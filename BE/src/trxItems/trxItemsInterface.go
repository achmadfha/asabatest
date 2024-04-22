package trxItems

import "BE-shop/models/dto/trxItemsDto"

type TrxItemsRepository interface {
	CreateTrxItems(trx trxItemsDto.TrxItems) error
	RetrieveAllTrxItems(transactionType string) (trxItems []trxItemsDto.TrxItemsRes, err error)
	DeleteTrxItems(trxID string) error
	UpdateTrxItemsByID(items trxItemsDto.TrxUpdateReq) error
	RetrieveTrxItemsByID(trxID string) (trxItemsDto.TrxItemsRes, error)
}

type TrxItemsUseCase interface {
	CreateTrxItems(trx trxItemsDto.TrxItemsReq) (trxItems trxItemsDto.TrxItemsRes, err error)
	RetrieveAllTrxItems(transactionType string) (trxItems []trxItemsDto.TrxItemsRes, err error)
	DeleteTrxItems(trxID string) error
	UpdateTrxItems(items trxItemsDto.TrxItemsUpdateReq) error
	RetrieveTrxItemsByID(trxID string) (trxItemsDto.TrxItemsRes, error)
}
