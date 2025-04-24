package controllers

import (
	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IncomeController interface {
	GetAllIncome(ctx *gin.Context)
	GetIncomeByInvoiceIdNumber(ctx *gin.Context)
	CreateIncome(ctx *gin.Context)
	UpdateIncome(ctx *gin.Context)
	DeleteIncome(ctx *gin.Context)
}

type incomeController struct {
	tokenService  services.TokenService
	incomeService services.IncomeService
}

func NewIncomeController(
	tokenService services.TokenService,
	incomeService services.IncomeService,
) IncomeController {
	return &incomeController{
		tokenService:  tokenService,
		incomeService: incomeService,
	}
}

func (c *incomeController) GetAllIncome(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	incomes, err := c.incomeService.GetAllIncome(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve income", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved income", incomes)
	ctx.JSON(http.StatusOK, res)
}

func (c *incomeController) GetIncomeByInvoiceIdNumber(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	incomeInvoiceIdNumber := ctx.Param("income_invoice_id_number")
	parsedIncomeInvoiceIdNumber, err := strconv.Atoi(incomeInvoiceIdNumber)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Income Invoice Id Number tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	income, err := c.incomeService.GetIncomeByInvoiceIdNumber(ctx.Request.Context(), parsedIncomeInvoiceIdNumber)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve income", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved income", income)
	ctx.JSON(http.StatusOK, res)
}

func (c *incomeController) CreateIncome(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.CreateIncomeRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	income, err := c.incomeService.CreateIncome(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to save income", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Data income successfully saved", income)
	ctx.JSON(http.StatusCreated, res)
}

func (c *incomeController) UpdateIncome(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.UpdateIncomeRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	incomeInvoiceIdNumber := ctx.Param("income_invoice_id_number")
	parsedIncomeInvoiceIdNumber, err := strconv.Atoi(incomeInvoiceIdNumber)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Income Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	income, err := c.incomeService.UpdateIncome(ctx, parsedIncomeInvoiceIdNumber, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update income", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Income successfully updated", income)
	ctx.JSON(http.StatusOK, res)
}

func (c *incomeController) DeleteIncome(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	incomeInvoiceIdNumber := ctx.Param("income_invoice_id_number")
	parsedIncomeInvoiceIdNumber, err := strconv.Atoi(incomeInvoiceIdNumber)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Income Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.incomeService.DeleteIncome(ctx, parsedIncomeInvoiceIdNumber); err != nil {
		res := utils.BuildResponseFailed("Failed to delete income", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Income successfully deleted", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
