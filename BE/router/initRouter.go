package router

import (
	"BE-shop/src/authentication/authenticationDelivery"
	"BE-shop/src/authentication/authenticationRepository"
	"BE-shop/src/authentication/authenticationUseCase"
	"BE-shop/src/items/itemsDelivery"
	"BE-shop/src/items/itemsRepository"
	"BE-shop/src/items/itemsUseCase"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	authenticationRepo := authenticationRepository.NewAuthenticationRepository(db)
	itemsRepo := itemsRepository.NewItemsRepository(db)

	// usecase
	authenticationUC := authenticationUseCase.NewAuthenticationUseCase(authenticationRepo)
	itemsUC := itemsUseCase.NewItemsUseCase(itemsRepo)

	// delivery
	authenticationDelivery.NewAuthenticationDelivery(v1Group, authenticationUC)
	itemsDelivery.NewItemsDelivery(v1Group, itemsUC)
}
