package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/mrbooi/event_backend/domain"
	"github.com/mrbooi/event_backend/utils"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			authToken := tokenArray[1]
			authorized, err := utils.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := utils.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
