package controllers

import (
	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentMethodController interface {
	GetAllPaymentMethod(ctx *gin.Context)
	GetPaymentMethodById(ctx *gin.Context)
	CreatePaymentMethod(ctx *gin.Context)
	UpdatePaymentMethod(ctx *gin.Context)
	DeletePaymentMethod(ctx *gin.Context)
}

type paymentMethodController struct {
	tokenService         services.TokenService
	paymentMethodService services.PaymentMethodService
}

func NewPaymentMethodController(
	tokenService services.TokenService,
	paymentMethodService services.PaymentMethodService,
) PaymentMethodController {
	return &paymentMethodController{
		tokenService:         tokenService,
		paymentMethodService: paymentMethodService,
	}
}

func (c *paymentMethodController) GetAllPaymentMethod(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	paymentMethods, err := c.paymentMethodService.GetAllPaymentMethod(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve payment method", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved payment method", paymentMethods)
	ctx.JSON(http.StatusOK, res)
}

func (c *paymentMethodController) GetPaymentMethodById(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	paymentMethodId := ctx.Param("payment_method_id")
	parsedPaymentMethodId, err := strconv.Atoi(paymentMethodId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Payment method Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	paymentMethod, err := c.paymentMethodService.GetPaymentMethodById(ctx.Request.Context(), parsedPaymentMethodId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve payment method", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved payment method", paymentMethod)
	ctx.JSON(http.StatusOK, res)
}

func (c *paymentMethodController) CreatePaymentMethod(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.PaymentMethodRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	paymentMethod, err := c.paymentMethodService.CreatePaymentMethod(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to save payment method", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Data payment method successfully saved", paymentMethod)
	ctx.JSON(http.StatusCreated, res)
}

func (c *paymentMethodController) UpdatePaymentMethod(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.PaymentMethodRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	paymentMethodId := ctx.Param("payment_method_id")
	parsedPaymentMethodId, err := strconv.Atoi(paymentMethodId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Payment Method Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	paymentMethod, err := c.paymentMethodService.UpdatePaymentMethod(ctx, parsedPaymentMethodId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update payment method", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Payment Method successfully updated", paymentMethod)
	ctx.JSON(http.StatusOK, res)
}

func (c *paymentMethodController) DeletePaymentMethod(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	paymentMethodId := ctx.Param("payment_method_id")
	parsedPaymentMethodId, err := strconv.Atoi(paymentMethodId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Payment method Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.paymentMethodService.DeletePaymentMethod(ctx, parsedPaymentMethodId); err != nil {
		res := utils.BuildResponseFailed("Failed to delete payment method", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Payment method successfully deleted", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
