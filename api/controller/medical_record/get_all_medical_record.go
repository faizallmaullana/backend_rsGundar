// GET GetAllMedicalRecord <= /medical_record/all

package medical_record

import (
	"net/http"
	"strings"

	"github.com/faizallmaullana/be_rsGundar/api/encryption"
	"github.com/faizallmaullana/be_rsGundar/api/models"
	"github.com/gin-gonic/gin"
)

func GetAllMedicalRecord(c *gin.Context) {
	var medicalRecord []models.MedicalRecord
	db := models.DB.Preload("Dokter.ProfileDokter.Poli").Preload("Pasien").Preload("Dokter.Profile").Preload("Diagnosis")
	if err := db.Find(&medicalRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database tidak bisa diakses"})
		return
	}

	// var decryptedData []map[string]interface{}
	// for _, record := range MedicalRecord {
	// 	// formated tanggal lahir
	// 	// Assuming you have a time.Time variable named "myDate"
	// 	myDate := record.CreatedAt

	// 	// Format the date to "02 01 06" layout
	// 	formattedDate := myDate.Format("02-01-2006")

	// 	namaPasien := strings.Title(encryption.Decrypt(record.Pasien.Nama))
	// 	diagnosis := strings.Title(encryption.Decrypt(record.Diagnosis.Diagnosis))
	// 	decryptedData = append(decryptedData, map[string]interface{}{
	// 		"id":                record.ID,
	// 		"nama_pasien":       namaPasien,
	// 		"tanggal_kunjungan": formattedDate,
	// 		"diagnosis":         diagnosis,
	// 	})
	// }

	var dataMedis []map[string]interface{}
	for _, record := range medicalRecord {

		// formated tanggal lahir
		// Assuming you have a time.Time variable named "myDate"
		myDate := record.CreatedAt

		// Format the date to "02 01 06" layout
		formattedDate := myDate.Format("02-01-2006")

		// dekripsi

		nama := titleCase(encryption.Decrypt(record.Pasien.Nama))
		dokter := titleCase(encryption.Decrypt(record.Dokter.Profile.Nama))
		poli := titleCase(encryption.Decrypt(record.Dokter.ProfileDokter.Poli.Poli))
		diagnosis := titleCase(encryption.Decrypt(record.Diagnosis.Diagnosis))

		dataMedis = append(dataMedis, map[string]interface{}{
			"id":            record.ID,
			"pasien":        nama,
			"poli":          poli,
			"diagnosis":     diagnosis,
			"dokter":        dokter,
			"biaya":         record.Biaya,
			"tanggal_lahir": formattedDate,
		})
	}

	c.JSON(http.StatusOK, gin.H{"medical_record": dataMedis})
}

func titleCase(data string) string {
	return strings.Title(data)
}
