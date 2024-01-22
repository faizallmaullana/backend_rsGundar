package medical_record

import (
	"net/http"
	"strings"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

// r.GET("/api/v1/resources/dokter/pemeriksaan/list/:idDokter", medical_record.GetDokterBaseID) // tested

func GetDokterBaseID(c *gin.Context) {
	var MedicalRecord []models.MedicalRecord
	db := models.DB.Where("id_dokter = ?", c.Param("idDokter"))
	dbPreload := db.Preload("Pasien").Preload("Diagnosis")
	if err := dbPreload.Find(&MedicalRecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "database tidak dapat diakses"})
		return
	}

	var decryptedData []map[string]interface{}
	for _, record := range MedicalRecord {
		// formated tanggal lahir
		// Assuming you have a time.Time variable named "myDate"
		myDate := record.CreatedAt

		// Format the date to "02 01 06" layout
		formattedDate := myDate.Format("02-01-2006")

		namaPasien := strings.Title(encryption.Decrypt(record.Pasien.Nama))
		diagnosis := strings.Title(encryption.Decrypt(record.Diagnosis.Diagnosis))
		decryptedData = append(decryptedData, map[string]interface{}{
			"id":                record.ID,
			"nama_pasien":       namaPasien,
			"tanggal_kunjungan": formattedDate,
			"diagnosis":         diagnosis,
		})
	}

	c.JSON(http.StatusOK, gin.H{"list": decryptedData})
}
