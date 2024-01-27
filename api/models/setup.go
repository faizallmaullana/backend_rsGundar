// here is setup for the db

package models

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	pwd, _ := os.Getwd()
	fmt.Println("Current working directory:", pwd)
	database, err := gorm.Open("sqlite3", "medicalRecord.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	// migration
	database.AutoMigrate(
		// profile
		&Users{},
		&Profile{},
		&ProfileDokter{},
		&DBToken{},

		// medical records
		&Pasien{},
		&Poli{},
		&Diagnosis{},
		&TempPendaftaran{},
		&MedicalRecord{},
		&Income{},
	)

	DB = database
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
