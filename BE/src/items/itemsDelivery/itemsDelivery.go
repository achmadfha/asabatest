package itemsDelivery

import (
	"BE-shop/models/constants"
	"BE-shop/models/dto/itemsDto"
	"BE-shop/models/dto/json"
	"BE-shop/pkg/validation"
	"BE-shop/src/items"
	"github.com/gin-gonic/gin"
)

type itemsDelivery struct {
	itemsUC items.ItemsUseCase
}

func NewItemsDelivery(v1Group *gin.RouterGroup, itemsUC items.ItemsUseCase) {
	handler := itemsDelivery{
		itemsUC: itemsUC,
	}

	itemsGroup := v1Group.Group("/items")
	{
		itemsGroup.POST("", handler.CreateItem)
		itemsGroup.GET("", handler.RetrieveAllItems)
	}
}

func (it itemsDelivery) CreateItem(ctx *gin.Context) {
	var req itemsDto.ItemsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	validationErr := validation.ValidationItems(req)
	if len(validationErr) > 0 {
		json.NewResponseBadRequest(ctx, validationErr, constants.BadReqMsg, constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	items, err := it.itemsUC.CreateItem(req)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, items, nil, "Success Create Items", constants.ServiceCodeItems, constants.SuccessCode)
}

func (it itemsDelivery) RetrieveAllItems(ctx *gin.Context) {
	transactionType := ctx.Query("transactionsType")
	itemsData, err := it.itemsUC.RetrieveAllItems(transactionType)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, itemsData, nil, "Success Retrieve All Items", constants.ServiceCodeItems, constants.SuccessCode)
}
