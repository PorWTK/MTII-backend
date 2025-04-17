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
	"os"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware sets the CORS headers and completes OPTIONS pre‑flight requests.
func CORSMiddleware() gin.HandlerFunc {
	// allowOrigin can be configured via env for flexibility
	allowOrigin := os.Getenv("FRONTEND_ORIGIN")
	if allowOrigin == "" {
		// fall back to your deployed frontend URL
		allowOrigin = "https://mtii-production.up.railway.app"
	}

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
