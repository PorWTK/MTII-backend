package controllers

import (
	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReceiverController interface {
	GetAllReceiver(ctx *gin.Context)
	GetReceiverById(ctx *gin.Context)
	CreateReceiver(ctx *gin.Context)
	UpdateReceiver(ctx *gin.Context)
	DeleteReceiver(ctx *gin.Context)
}

type receiverController struct {
	tokenService    services.TokenService
	receiverService services.ReceiverService
}

func NewReceiverController(
	tokenService services.TokenService,
	receiverService services.ReceiverService,
) ReceiverController {
	return &receiverController{
		tokenService:    tokenService,
		receiverService: receiverService,
	}
}

func (c *receiverController) GetAllReceiver(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	receivers, err := c.receiverService.GetAllReceiver(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve receiver", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved receiver", receivers)
	ctx.JSON(http.StatusOK, res)
}

func (c *receiverController) GetReceiverById(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	receiverId := ctx.Param("receiver_id")
	parsedReceiverId, err := strconv.Atoi(receiverId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Receiver Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	receiver, err := c.receiverService.GetReceiverById(ctx.Request.Context(), parsedReceiverId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve receiver", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved receiver", receiver)
	ctx.JSON(http.StatusOK, res)
}

func (c *receiverController) CreateReceiver(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.CreateReceiverRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	receiver, err := c.receiverService.CreateReceiver(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to save receiver", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Data receiver successfully saved", receiver)
	ctx.JSON(http.StatusCreated, res)
}

func (c *receiverController) UpdateReceiver(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var req dtos.UpdateReceiverRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	receiverId := ctx.Param("receiver_id")
	parsedReceiverId, err := strconv.Atoi(receiverId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Receiver Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	receiver, err := c.receiverService.UpdateReceiver(ctx, parsedReceiverId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update receiver", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Receiver successfully updated", receiver)
	ctx.JSON(http.StatusOK, res)
}

func (c *receiverController) DeleteReceiver(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	_, err := c.tokenService.GetUserIdByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Invalid token", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	receiverId := ctx.Param("receiver_id")
	parsedReceiverId, err := strconv.Atoi(receiverId)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to process the request", "Receiver Id tidak valid", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.receiverService.DeleteReceiver(ctx, parsedReceiverId); err != nil {
		res := utils.BuildResponseFailed("Failed to delete receiver", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Receiver successfully deleted", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
