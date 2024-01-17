package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/faizallmaullana/be_rsGundar/models"
)

// initilaize the cors middleware
var corsConfig = cors.DefaultConfig()

func init() {
	// allow all origins
	corsConfig.AllowAllOrigins = true
}

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	// connect to database
	models.ConnectDatabase()
	r.Use(cors.New(corsConfig))

	// ROUTES

	// run the server
	r.Run(":9888")
}
