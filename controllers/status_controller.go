package controllers

import (
	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatusController interface {
	GetAllStatus(ctx *gin.Context)
	GetStatusById(ctx *gin.Context)
	CreateStatus(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	DeleteStatus(ctx *gin.Context)
}

type statusController struct {
	tokenService  services.TokenService
	statusService services.StatusService
}

func NewStatusController(
	tokenService services.TokenService,
	statusService services.StatusService,
) StatusController {
	return &statusController{
		tokenService:  tokenService,
		statusService: statusService,
	}
}

func (c *statusController) GetAllStatus(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	statuses, err := c.statusService.GetAllStatus(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal mendapatkan status", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil mendapatkan status", statuses)
	ctx.JSON(http.StatusOK, res)
}

func (c *statusController) GetStatusById(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	statusId := ctx.Param("status_id")
	parsedStatusId, err := strconv.Atoi(statusId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Status Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	status, err := c.statusService.GetStatusById(ctx.Request.Context(), parsedStatusId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal mendapatkan status", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil mendapatkan status", status)
	ctx.JSON(http.StatusOK, res)
}

func (c *statusController) CreateStatus(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.StatusRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Gagal mendapatkan request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	status, err := c.statusService.CreateStatus(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal menyimpan status", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Data status berhasil disimpan", status)
	ctx.JSON(http.StatusCreated, res)
}

func (c *statusController) UpdateStatus(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.StatusRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Gagal mendapatkan request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	statusId := ctx.Param("status_id")
	parsedStatusId, err := strconv.Atoi(statusId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Status Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	status, err := c.statusService.UpdateStatus(ctx, parsedStatusId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal mengupdate status", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Status berhasil diupdate", status)
	ctx.JSON(http.StatusOK, res)
}

func (c *statusController) DeleteStatus(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Token tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	statusId := ctx.Param("status_id")
	parsedStatusId, err := strconv.Atoi(statusId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal memproses request", "Status Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.statusService.DeleteStatus(ctx, parsedStatusId); err != nil {
		res := utils.BuildResponseFailed("Gagal menghapus status", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Status berhasil dihapus", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
