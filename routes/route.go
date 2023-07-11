package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrbooi/event_backend/bootstrap"
	"github.com/mrbooi/event_backend/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	v1 := gin.Group("/v1")

	// health route
	healthCheckRouter(env, timeout, v1)

}
