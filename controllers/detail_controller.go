package controllers

import (
	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DetailController interface {
	GetAllDetail(ctx *gin.Context)
	GetDetailById(ctx *gin.Context)
	CreateDetail(ctx *gin.Context)
	UpdateDetail(ctx *gin.Context)
	DeleteDetail(ctx *gin.Context)
}

type detailController struct {
	tokenService  services.TokenService
	detailService services.DetailService
}

func NewDetailController(
	tokenService services.TokenService,
	detailService services.DetailService,
) DetailController {
	return &detailController{
		tokenService:  tokenService,
		detailService: detailService,
	}
}

func (c *detailController) GetAllDetail(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	details, err := c.detailService.GetAllDetail(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve detail", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved detail", details)
	ctx.JSON(http.StatusOK, res)
}

func (c *detailController) GetDetailById(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	detailId := ctx.Param("detail_id")
	parsedDetailId, err := strconv.Atoi(detailId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Detail Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	detail, err := c.detailService.GetDetailById(ctx.Request.Context(), parsedDetailId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve detail", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved detail", detail)
	ctx.JSON(http.StatusOK, res)
}

func (c *detailController) CreateDetail(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.CreateDetailRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	detail, err := c.detailService.CreateDetail(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to save detail", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Data detail successfully saved", detail)
	ctx.JSON(http.StatusCreated, res)
}

func (c *detailController) UpdateDetail(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.UpdateDetailRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	detailId := ctx.Param("detail_id")
	parsedDetailId, err := strconv.Atoi(detailId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Detail Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	detail, err := c.detailService.UpdateDetail(ctx, parsedDetailId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update detail", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Detail successfully updated", detail)
	ctx.JSON(http.StatusOK, res)
}

func (c *detailController) DeleteDetail(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	detailId := ctx.Param("detail_id")
	parsedDetailId, err := strconv.Atoi(detailId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Detail Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.detailService.DeleteDetail(ctx, parsedDetailId); err != nil {
		res := utils.BuildResponseFailed("Failed to delete detail", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Detail successfully deleted", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
