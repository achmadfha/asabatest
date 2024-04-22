package trxItemsDelivery

import (
	"BE-shop/models/constants"
	"BE-shop/models/dto/json"
	"BE-shop/models/dto/trxItemsDto"
	"BE-shop/pkg/validation"
	"BE-shop/src/trxItems"
	"github.com/gin-gonic/gin"
)

type trxItemsDelivery struct {
	trxItemsUC trxItems.TrxItemsUseCase
}

func NewTrxItemsDelivery(v1Group *gin.RouterGroup, trxItemsUC trxItems.TrxItemsUseCase) {
	handler := trxItemsDelivery{
		trxItemsUC,
	}

	trxItemsGroup := v1Group.Group("/transaction")
	{
		trxItemsGroup.POST("", handler.CreateTrxItems)
		trxItemsGroup.GET("", handler.RetrieveAllTrxItems)
		trxItemsGroup.DELETE("/:id", handler.DeleteTrxItems)
	}
}

func (t trxItemsDelivery) CreateTrxItems(ctx *gin.Context) {
	var req trxItemsDto.TrxItemsReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeTrx, constants.GeneralErrCode)
		return
	}

	validationErr := validation.TrxItemsValidation(req)
	if len(validationErr) > 0 {
		json.NewResponseBadRequest(ctx, validationErr, constants.BadReqMsg, constants.ServiceCodeTrx, constants.GeneralErrCode)
		return
	}

	trxItems, err := t.trxItemsUC.CreateTrxItems(req)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "items doesn't exist", constants.ServiceCodeTrx, constants.Forbidden)
			return
		}
		if err.Error() == "02" {
			json.NewResponseForbidden(ctx, "amount is not enough", constants.ServiceCodeTrx, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeTrx, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, trxItems, nil, "Success Create Transactions Items", constants.ServiceCodeTrx, constants.SuccessCode)
}

func (t trxItemsDelivery) RetrieveAllTrxItems(ctx *gin.Context) {
	transactionType := ctx.Query("transactionsType")

	trxItems, err := t.trxItemsUC.RetrieveAllTrxItems(transactionType)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeTrx, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, trxItems, nil, "Success Retrieve All Transactions Items", constants.ServiceCodeTrx, constants.SuccessCode)
}

func (t trxItemsDelivery) DeleteTrxItems(ctx *gin.Context) {
	trxID := ctx.Param("id")

	err := t.trxItemsUC.DeleteTrxItems(trxID)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseForbidden(ctx, "transactions items doesn't exist", constants.ServiceCodeTrx, constants.Forbidden)
			return
		}
		json.NewResponseError(ctx, err.Error(), constants.ServiceCodeTrx, constants.GeneralErrCode)
		return
	}

	json.NewResponseSuccess(ctx, nil, nil, "Success Delete Transactions Items", constants.ServiceCodeTrx, constants.SuccessCode)
}
