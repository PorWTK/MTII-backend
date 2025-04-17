// package middlewares

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Allow-Credentials", "true")
// 		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

// 		if c.Request.Method == http.MethodOptions {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware sets CORS headers and handles preflight requests.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: replace this with your actual frontend URL or read from env
		allowedOrigin := "https://mtii-production.up.railway.app"

		c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		// Respond OK to preflight requests
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
