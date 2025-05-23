package controllers

import (
	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlatformController interface {
	GetAllPlatform(ctx *gin.Context)
	GetPlatformById(ctx *gin.Context)
	CreatePlatform(ctx *gin.Context)
	UpdatePlatform(ctx *gin.Context)
	DeletePlatform(ctx *gin.Context)
}

type platformController struct {
	tokenService    services.TokenService
	platformService services.PlatformService
}

func NewPlatformController(
	tokenService services.TokenService,
	platformService services.PlatformService,
) PlatformController {
	return &platformController{
		tokenService:    tokenService,
		platformService: platformService,
	}
}

func (c *platformController) GetAllPlatform(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	platforms, err := c.platformService.GetAllPlatform(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve platform", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved platform", platforms)
	ctx.JSON(http.StatusOK, res)
}

func (c *platformController) GetPlatformById(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	platformId := ctx.Param("platform_id")
	parsedPlatformId, err := strconv.Atoi(platformId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Platform Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	platform, err := c.platformService.GetPlatformById(ctx.Request.Context(), parsedPlatformId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve platform", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved platform", platform)
	ctx.JSON(http.StatusOK, res)
}

func (c *platformController) CreatePlatform(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.PlatformRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	platform, err := c.platformService.CreatePlatform(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to save platform", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Data platform successfully saved", platform)
	ctx.JSON(http.StatusCreated, res)
}

func (c *platformController) UpdatePlatform(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.PlatformRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	platformId := ctx.Param("platform_id")
	parsedPlatformId, err := strconv.Atoi(platformId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Platform Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	platform, err := c.platformService.UpdatePlatform(ctx, parsedPlatformId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update platform", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Platform successfully updated", platform)
	ctx.JSON(http.StatusOK, res)
}

func (c *platformController) DeletePlatform(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	platformId := ctx.Param("platform_id")
	parsedPlatformId, err := strconv.Atoi(platformId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Platform Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.platformService.DeletePlatform(ctx, parsedPlatformId); err != nil {
		res := utils.BuildResponseFailed("Failed to delete platform", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Platform successfully deleted", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
