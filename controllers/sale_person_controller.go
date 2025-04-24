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
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	salePeople, err := c.salePersonService.GetAllSalePerson(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve sale person", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved sale person", salePeople)
	ctx.JSON(http.StatusOK, res)
}

func (c *salePersonController) GetSalePersonById(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	salePersonId := ctx.Param("sale_person_id")
	parsedSalePersonId, err := strconv.Atoi(salePersonId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Sale person Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	salePerson, err := c.salePersonService.GetSalePersonById(ctx.Request.Context(), parsedSalePersonId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve sale person", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved sale person", salePerson)
	ctx.JSON(http.StatusOK, res)
}

func (c *salePersonController) CreateSalePerson(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.SalePersonRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	salePerson, err := c.salePersonService.CreateSalePerson(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to save sale person", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Data sale person successfully saved", salePerson)
	ctx.JSON(http.StatusCreated, res)
}

func (c *salePersonController) UpdateSalePerson(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.SalePersonRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	salePersonId := ctx.Param("sale_person_id")
	parsedSalePersonId, err := strconv.Atoi(salePersonId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Sale Person Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	salePerson, err := c.salePersonService.UpdateSalePerson(ctx, parsedSalePersonId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update sale person", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Sale Person successfully updated", salePerson)
	ctx.JSON(http.StatusOK, res)
}

func (c *salePersonController) DeleteSalePerson(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	salePersonId := ctx.Param("sale_person_id")
	parsedSalePersonId, err := strconv.Atoi(salePersonId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Sale person Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.salePersonService.DeleteSalePerson(ctx, parsedSalePersonId); err != nil {
		res := utils.BuildResponseFailed("Failed to delete sale person", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Sale person successfully deleted", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
