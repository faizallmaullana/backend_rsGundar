// GET Statistik <= /api/v1/resources/statistik

package statistik

import (
	"net/http"

	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

func Statistik(c *gin.Context) {
	// deklarasi kebutuhan variabel
	var countMedicalRecords int32
	var countDiagnosis int32
	var countPasien int32

	// menghitung data dari database
	// admin
	var medicalRecord []models.MedicalRecord
	if err := models.DB.Find(&medicalRecord).Count(&countMedicalRecords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak bisa menghitung rekam medis"})
		return
	}

	var diagnosis []models.Diagnosis
	if err := models.DB.Find(&diagnosis).Count(&countDiagnosis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak bisa menghitung diagnosis"})
		return
	}

	var income models.Income
	if err := models.DB.First(&income).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak bisa menghitung income"})
		return
	}

	var Pasien []models.Pasien
	if err := models.DB.Find(&Pasien).Count(&countPasien).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak bisa menghitung pasien"})
		return
	}

	// =============================================

	// return data
	c.JSON(http.StatusOK, gin.H{
		"total_rekam_medis": countMedicalRecords,
		"total_diagnosis":   countDiagnosis,
		"total_income":      income.Income,
		"total_pasien":      countPasien,
	})
}
