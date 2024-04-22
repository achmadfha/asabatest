package itemsDelivery

import (
	"BE-shop/models/constants"
	"BE-shop/models/dto/itemsDto"
	"BE-shop/models/dto/json"
	"BE-shop/pkg/middleware"
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
		itemsGroup.POST("", middleware.JWTAuth("ADMIN"), handler.CreateItem)
		itemsGroup.GET("", middleware.JWTAuth("ADMIN"), handler.RetrieveAllItems)
		itemsGroup.GET("/:code", middleware.JWTAuth("ADMIN"), handler.RetrieveItemsByCode)
		itemsGroup.PUT("/:code", middleware.JWTAuth("ADMIN"), handler.UpdateItemsByCode)
		itemsGroup.DELETE("/:code", middleware.JWTAuth("ADMIN"), handler.DeleteItemsByCode)
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
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "items already exist", constants.ServiceCodeItems, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, items, nil, "Success Create Items", constants.ServiceCodeItems, constants.SuccessCode)
}

func (it itemsDelivery) RetrieveAllItems(ctx *gin.Context) {
	itemsData, err := it.itemsUC.RetrieveAllItems()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, itemsData, nil, "Success Retrieve All Items", constants.ServiceCodeItems, constants.SuccessCode)
}

func (it itemsDelivery) RetrieveItemsByCode(ctx *gin.Context) {
	code := ctx.Param("code")

	itemsData, err := it.itemsUC.RetrieveItemsByCode(code)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "items doesn't exist", constants.ServiceCodeItems, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, itemsData, nil, "Success Retrieve Items By Code", constants.ServiceCodeItems, constants.SuccessCode)
}

func (it itemsDelivery) UpdateItemsByCode(ctx *gin.Context) {
	var req itemsDto.ItemsUpdate
	code := ctx.Param("code")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	itemsData := itemsDto.ItemsUpdate{
		Code:            code,
		Amount:          req.Amount,
		StatusActive:    req.StatusActive,
		TransactionType: req.TransactionType,
	}

	validationErr := validation.ValidationItemsUpdate(itemsData)
	if len(validationErr) > 0 {
		json.NewResponseBadRequest(ctx, validationErr, constants.BadReqMsg, constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	data, err := it.itemsUC.UpdateItemsByCode(itemsData)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "items doesn't exist", constants.ServiceCodeItems, constants.Forbidden)
			return
		}
		if err.Error() == "02" {
			json.NewResponseForbidden(ctx, "amount is not enough", constants.ServiceCodeItems, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, data, nil, "Success Update Items By Code", constants.ServiceCodeItems, constants.SuccessCode)
}

func (it itemsDelivery) DeleteItemsByCode(ctx *gin.Context) {
	code := ctx.Param("code")

	err := it.itemsUC.DeleteItemsByCode(code)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "items doesn't exist", constants.ServiceCodeItems, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeItems, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, nil, nil, "Success Delete Items By Code", constants.ServiceCodeItems, constants.SuccessCode)
}
