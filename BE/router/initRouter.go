package router

import (
	"BE-shop/src/authentication/authenticationDelivery"
	"BE-shop/src/authentication/authenticationRepository"
	"BE-shop/src/authentication/authenticationUseCase"
	"BE-shop/src/items/itemsDelivery"
	"BE-shop/src/items/itemsRepository"
	"BE-shop/src/items/itemsUseCase"
	"BE-shop/src/trxItems/trxItemsDelivery"
	"BE-shop/src/trxItems/trxItemsRepository"
	"BE-shop/src/trxItems/trxItemsUseCase"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	authenticationRepo := authenticationRepository.NewAuthenticationRepository(db)
	itemsRepo := itemsRepository.NewItemsRepository(db)
	trxItemsRepo := trxItemsRepository.NewTrxItemsRepository(db)

	// usecase
	authenticationUC := authenticationUseCase.NewAuthenticationUseCase(authenticationRepo)
	itemsUC := itemsUseCase.NewItemsUseCase(itemsRepo)
	trxItemsUC := trxItemsUseCase.NewTrxItemsUseCase(trxItemsRepo, itemsRepo)

	// delivery
	authenticationDelivery.NewAuthenticationDelivery(v1Group, authenticationUC)
	itemsDelivery.NewItemsDelivery(v1Group, itemsUC)
	trxItemsDelivery.NewTrxItemsDelivery(v1Group, trxItemsUC)
}
