package controllers

import (
	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BankController interface {
	GetAllBank(ctx *gin.Context)
	GetBankById(ctx *gin.Context)
	CreateBank(ctx *gin.Context)
	UpdateBank(ctx *gin.Context)
	DeleteBank(ctx *gin.Context)
}

type bankController struct {
	tokenService services.TokenService
	bankService  services.BankService
}

func NewBankController(
	tokenService services.TokenService,
	bankService services.BankService,
) BankController {
	return &bankController{
		tokenService: tokenService,
		bankService:  bankService,
	}
}

func (c *bankController) GetAllBank(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	banks, err := c.bankService.GetAllBank(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal mendapatkan bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil mendapatkan bank", banks)
	ctx.JSON(http.StatusOK, res)
}

func (c *bankController) GetBankById(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	bankId := ctx.Param("bank_id")
	parsedBankId, err := strconv.Atoi(bankId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Bank Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	bank, err := c.bankService.GetBankById(ctx.Request.Context(), parsedBankId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal mendapatkan bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil mendapatkan bank", bank)
	ctx.JSON(http.StatusOK, res)
}

func (c *bankController) CreateBank(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.BankRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Gagal mendapatkan request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	bank, err := c.bankService.CreateBank(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal menyimpan bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Data bank berhasil disimpan", bank)
	ctx.JSON(http.StatusCreated, res)
}

func (c *bankController) UpdateBank(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.BankRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Gagal mendapatkan request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	bankId := ctx.Param("bank_id")
	parsedBankId, err := strconv.Atoi(bankId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Bank Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	bank, err := c.bankService.UpdateBank(ctx, parsedBankId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal mengupdate bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Bank berhasil diupdate", bank)
	ctx.JSON(http.StatusOK, res)
}

func (c *bankController) DeleteBank(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	bankId := ctx.Param("bank_id")
	parsedBankId, err := strconv.Atoi(bankId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Bank Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.bankService.DeleteBank(ctx, parsedBankId); err != nil {
		res := utils.BuildResponseFailed("Gagal menghapus bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Bank berhasil dihapus", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
