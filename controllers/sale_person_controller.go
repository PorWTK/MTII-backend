package controllers

import (
	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SalePersonController interface {
	GetAllSalePerson(ctx *gin.Context)
	GetSalePersonById(ctx *gin.Context)
	CreateSalePerson(ctx *gin.Context)
	UpdateSalePerson(ctx *gin.Context)
	DeleteSalePerson(ctx *gin.Context)
}

type salePersonController struct {
	tokenService      services.TokenService
	salePersonService services.SalePersonService
}

func NewSalePersonController(
	tokenService services.TokenService,
	salePersonService services.SalePersonService,
) SalePersonController {
	return &salePersonController{
		tokenService:      tokenService,
		salePersonService: salePersonService,
	}
}

func (c *salePersonController) GetAllSalePerson(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	salePeople, err := c.salePersonService.GetAllSalePerson(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal mendapatkan sale person", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil mendapatkan sale person", salePeople)
	ctx.JSON(http.StatusOK, res)
}

func (c *salePersonController) GetSalePersonById(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	salePersonId := ctx.Param("sale_person_id")
	parsedSalePersonId, err := strconv.Atoi(salePersonId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Sale person Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	salePerson, err := c.salePersonService.GetSalePersonById(ctx.Request.Context(), parsedSalePersonId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal mendapatkan sale person", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil mendapatkan sale person", salePerson)
	ctx.JSON(http.StatusOK, res)
}

func (c *salePersonController) CreateSalePerson(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.SalePersonRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Gagal mendapatkan request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	salePerson, err := c.salePersonService.CreateSalePerson(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal menyimpan sale person", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Data sale person berhasil disimpan", salePerson)
	ctx.JSON(http.StatusCreated, res)
}

func (c *salePersonController) UpdateSalePerson(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.SalePersonRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Gagal mendapatkan request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	salePersonId := ctx.Param("sale_person_id")
	parsedSalePersonId, err := strconv.Atoi(salePersonId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Sale Person Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	salePerson, err := c.salePersonService.UpdateSalePerson(ctx, parsedSalePersonId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal mengupdate sale person", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Sale Person berhasil diupdate", salePerson)
	ctx.JSON(http.StatusOK, res)
}

func (c *salePersonController) DeleteSalePerson(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	salePersonId := ctx.Param("sale_person_id")
	parsedSalePersonId, err := strconv.Atoi(salePersonId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Sale person Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.salePersonService.DeleteSalePerson(ctx, parsedSalePersonId); err != nil {
		res := utils.BuildResponseFailed("Gagal menghapus sale person", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Sale person berhasil dihapus", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
