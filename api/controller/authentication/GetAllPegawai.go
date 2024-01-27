package authentication

import (
	"net/http"

	"github.com/faizallmaullana/be_rsGundar/api/encryption"
	"github.com/faizallmaullana/be_rsGundar/api/models"
	"github.com/gin-gonic/gin"
)

func GetAllPegawai(c *gin.Context) {
	var users []models.Users
	dbPreload := models.DB.Preload("Profile")
	if err := dbPreload.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var dataKaryawan []map[string]interface{}
	for _, record := range users {
		// formated tanggal lahir
		// Assuming you have a time.Time variable named "myDate"
		myDate := record.CreatedAt

		// Format the date to "02 01 06" layout
		formattedDate := myDate.Format("02-01-2006")

		// dektipsi
		nama := titleCase(encryption.Decrypt(record.Profile.Nama))
		posisi := titleCase(encryption.Decrypt(record.Role))

		dataKaryawan = append(dataKaryawan, map[string]interface{}{
			"id":            record.ID,
			"nama":          nama,
			"tanggal_masuk": formattedDate,
			"posisi":        posisi,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"dataKaryawan": dataKaryawan,
	})
}
