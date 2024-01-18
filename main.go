package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/faizallmaullana/be_rsGundar/controller/authentication"
	"github.com/faizallmaullana/be_rsGundar/controller/medical_record"
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

	// authentication
	r.POST("/api/v1/resources/registration/dokter", authentication.RegistrasiDokter)
	r.POST("/api/v1/resources/registration/admin", authentication.Registrasi)
	r.POST("/api/v1/resources/registration/staffPendaftaran", authentication.RegistrasiStaffPendaftaran)
	r.POST("/api/v1/resources/login", authentication.Login)

	// poli
	r.POST("/api/v1/resources/poli", medical_record.AddPoli)

	// run the server
	r.Run(":3200")
}
