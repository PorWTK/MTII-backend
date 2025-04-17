package controllers

import (
	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	LoginUser(ctx *gin.Context)
	LogoutUser(ctx *gin.Context)
}

type userController struct {
	tokenService services.TokenService
	userService  services.UserService
}

func NewUserController(
	tokenService services.TokenService,
	userService services.UserService,
) UserController {
	return &userController{
		tokenService: tokenService,
		userService:  userService,
	}
}

func (c *userController) LoginUser(ctx *gin.Context) {
	var LoginRequest dtos.LoginRequest
	err := ctx.ShouldBind(&LoginRequest)
	if err != nil {
		response := utils.BuildResponseFailed("Gagal mendapatkan request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.userService.VerifyCredential(ctx.Request.Context(), LoginRequest)
	if err != nil {
		response := utils.BuildResponseFailed("User gagal login", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.BuildResponseSuccess("User berhasil login", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) LogoutUser(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		res := utils.BuildResponseFailed("User gagal logout", "Token tidak ada", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	err := c.tokenService.InvalidateToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("User gagal logut", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctx.Header("Set-Cookie", "token=; Path=/; Max-Age=-1")
	ctx.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")

	res := utils.BuildResponseSuccess("User berhasil logout", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
