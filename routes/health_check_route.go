package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrbooi/event_backend/bootstrap"
)

func healthCheckRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	group.GET("/healthcheck", func(c *gin.Context) {
		c.String(200, "Success")
	})
}
