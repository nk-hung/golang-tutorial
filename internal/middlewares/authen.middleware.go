package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/nk-hung/go-ecommerce-backend-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
		if token != "invalid-token" {
			response.ErrorResponse(c, response.ErrodeTokenInvalid, "Token is invalid")
			c.Abort()
			return
		}
		// Do something here
		c.Next()
	}
}
