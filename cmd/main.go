package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mrbooi/event_backend/bootstrap"
	routes "github.com/mrbooi/event_backend/routes"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout)

	gin := gin.Default()

	routes.Setup(env, timeout, db, gin)

	// getup middleware to deserializeUser
	// gin.Use(deserializeUser)

	err := gin.Run(env.ServerAddress)
	if err != nil {
		fmt.Println("Failed while setting up App routes.")
	}

}
