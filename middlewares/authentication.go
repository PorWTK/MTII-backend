package middlewares

import (
	"mtii-backend/services"
	"mtii-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(tokenService services.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildResponseFailed("Failed to process the Request", "Token Not Found", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.BuildResponseFailed("Failed to process the Request", "Invalid token", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := tokenService.ValidateToken(authHeader)
		if err != nil {
			response := utils.BuildResponseFailed("Failed to process the Request", "Invalid token", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !token.Valid {
			response := utils.BuildResponseFailed("Failed to process the Request", "Access Denied", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userId, err := tokenService.GetUserIdByToken(authHeader)
		if err != nil {
			response := utils.BuildResponseFailed("Failed to process the Request", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Set("token", authHeader)
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
