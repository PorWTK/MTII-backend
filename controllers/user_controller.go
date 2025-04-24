// package controllers

// import (
// 	"mtii-backend/dtos"
// 	"mtii-backend/services"
// 	"mtii-backend/utils"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// type UserController interface {
// 	LoginUser(ctx *gin.Context)
// 	LogoutUser(ctx *gin.Context)
// }

// type userController struct {
// 	tokenService services.TokenService
// 	userService  services.UserService
// }

// func NewUserController(
// 	tokenService services.TokenService,
// 	userService services.UserService,
// ) UserController {
// 	return &userController{
// 		tokenService: tokenService,
// 		userService:  userService,
// 	}
// }

// func (c *userController) LoginUser(ctx *gin.Context) {
// 	var LoginRequest dtos.LoginRequest
// 	err := ctx.ShouldBind(&LoginRequest)
// 	if err != nil {
// 		response := utils.BuildResponseFailed("Failed to retrieve request", err.Error(), utils.EmptyObj{})
// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	res, err := c.userService.VerifyCredential(ctx.Request.Context(), LoginRequest)
// 	if err != nil {
// 		response := utils.BuildResponseFailed("User gagal login", err.Error(), utils.EmptyObj{})
// 		ctx.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := utils.BuildResponseSuccess("User berhasil login", res)
// 	ctx.JSON(http.StatusOK, response)
// }

// func (c *userController) LogoutUser(ctx *gin.Context) {
// 	token := ctx.GetHeader("Authorization")
// 	if token == "" {
// 		res := utils.BuildResponseFailed("User gagal logout", "Token tidak ada", utils.EmptyObj{})
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	token = strings.TrimPrefix(token, "Bearer ")

// 	err := c.tokenService.InvalidateToken(token)
// 	if err != nil {
// 		res := utils.BuildResponseFailed("User gagal logut", err.Error(), utils.EmptyObj{})
// 		ctx.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	ctx.Header("Set-Cookie", "token=; Path=/; Max-Age=-1")
// 	ctx.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")

// 	res := utils.BuildResponseSuccess("User berhasil logout", utils.EmptyObj{})
// 	ctx.JSON(http.StatusOK, res)
// }

package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"mtii-backend/dtos"
	"mtii-backend/services"
	"mtii-backend/utils"

	"github.com/gin-gonic/gin"
)

/* ────────────────────────────────────────────────────────── */
/* Interfaces & constructor                                  */
/* ────────────────────────────────────────────────────────── */

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

/* ────────────────────────────────────────────────────────── */
/* Login                                                     */
/* ────────────────────────────────────────────────────────── */

func (c *userController) LoginUser(ctx *gin.Context) {
	var req dtos.LoginRequest

	// 1) Bind ONLY JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,
			utils.BuildResponseFailed("Bad JSON", err.Error(), utils.EmptyObj{}))
		return
	}

	// 2) DEBUG: print exactly what arrived
	fmt.Printf("LOGIN DEBUG %#v\n", req)
	//     Example output:
	//     LOGIN DEBUG dtos.LoginRequest{Username:"user1", Password:"password "}

	// 3) Verify credentials
	res, err := c.userService.VerifyCredential(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusForbidden,
			utils.BuildResponseFailed("nvalid credentials", err.Error(), utils.EmptyObj{}))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("Login OK", res))
}

/* ────────────────────────────────────────────────────────── */
/* Logout                                                    */
/* ────────────────────────────────────────────────────────── */

func (c *userController) LogoutUser(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusBadRequest,
			utils.BuildResponseFailed("User gagal logout", "Token tidak ada", utils.EmptyObj{}))
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	if err := c.tokenService.InvalidateToken(token); err != nil {
		ctx.JSON(http.StatusBadRequest,
			utils.BuildResponseFailed("User gagal logout", err.Error(), utils.EmptyObj{}))
		return
	}

	// Clear cookie (if you set one)
	ctx.Header("Set-Cookie", "token=; Path=/; Max-Age=-1")
	ctx.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess("User berhasil logout", utils.EmptyObj{}))
}
