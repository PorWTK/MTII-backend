package controllers

import (
	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChannelController interface {
	GetAllChannel(ctx *gin.Context)
	GetChannelById(ctx *gin.Context)
	CreateChannel(ctx *gin.Context)
	UpdateChannel(ctx *gin.Context)
	DeleteChannel(ctx *gin.Context)
}

type channelController struct {
	tokenService   services.TokenService
	channelService services.ChannelService
}

func NewChannelController(
	tokenService services.TokenService,
	channelService services.ChannelService,
) ChannelController {
	return &channelController{
		tokenService:   tokenService,
		channelService: channelService,
	}
}

func (c *channelController) GetAllChannel(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	channels, err := c.channelService.GetAllChannel(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve channel", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved channel", channels)
	ctx.JSON(http.StatusOK, res)
}

func (c *channelController) GetChannelById(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	channelId := ctx.Param("channel_id")
	parsedChannelId, err := strconv.Atoi(channelId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Channel Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	channel, err := c.channelService.GetChannelById(ctx.Request.Context(), parsedChannelId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve channel", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved channel", channel)
	ctx.JSON(http.StatusOK, res)
}

func (c *channelController) CreateChannel(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.ChannelRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	channel, err := c.channelService.CreateChannel(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to save channel", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Data channel successfully saved", channel)
	ctx.JSON(http.StatusCreated, res)
}

func (c *channelController) UpdateChannel(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.ChannelRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	channelId := ctx.Param("channel_id")
	parsedChannelId, err := strconv.Atoi(channelId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Channel Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	channel, err := c.channelService.UpdateChannel(ctx, parsedChannelId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update channel", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Channel successfully updated", channel)
	ctx.JSON(http.StatusOK, res)
}

func (c *channelController) DeleteChannel(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	channelId := ctx.Param("channel_id")
	parsedChannelId, err := strconv.Atoi(channelId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Channel Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.channelService.DeleteChannel(ctx, parsedChannelId); err != nil {
		res := utils.BuildResponseFailed("Failed to delete channel", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Channel successfully deleted", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
