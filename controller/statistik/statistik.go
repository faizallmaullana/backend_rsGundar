// GET Statistik <= /api/v1/resources/statistik

package statistik

import (
	"net/http"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

func Statistik(c *gin.Context) {
	// deklarasi kebutuhan variabel
	var countMedicalRecords int64
	var countDiagnosis int64
	var countPasien int64

	var totalDokter int32
	var totalAdmin int32
	var totalStaff int32

	// menghitung data dari database
	// admin
	var User models.Users
	admin := encryption.Encrypt("admin")
	if err := models.DB.Where("role = ?", admin).Find(&User).Count(&totalAdmin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak bisa menghitung dokter"})
		return
	}

	// dokter
	dokter := encryption.Encrypt("dokter")
	if err := models.DB.Where("role = ?", dokter).Find(&User).Count(&totalDokter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak bisa menghitung dokter"})
		return
	}

	// staff
	staff := encryption.Encrypt("staffPendaftaran")
	if err := models.DB.Where("role = ?", staff).Find(&User).Count(&totalStaff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak bisa menghitung dokter"})
		return
	}

	// pasien
	var Pasien models.Pasien
	if err := models.DB.Find(&Pasien).Count(&countPasien).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data pasien tidak bisa diakses"})
		return
	}
	// diagnosis
	var Diagnosis models.Diagnosis
	if err := models.DB.Find(&Diagnosis).Count(&countDiagnosis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data diganosis tidak bisa diakses"})
		return
	}

	// income
	var Income models.Income
	if err := models.DB.First(&Income).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data income tidak bisa diakses"})
		return
	}

	// medical record
	var MedicalRecords models.MedicalRecord
	if err := models.DB.Find(&MedicalRecords).Count(&countMedicalRecords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data medical record tidak bisa diakses"})
		return
	}

	// poli
	var Poli models.Poli
	if err := models.DB.Find(&Poli).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data poli tidak bisa diakses"})
		return
	}

	// =============================================

	// return data
	c.JSON(http.StatusOK, gin.H{
		"total_rekam_medis": countMedicalRecords,
		"total_diagnosis":   countDiagnosis,
		"total_pasien":      countPasien,
		"total_income":      Income.Income,
		"total_admin":       totalAdmin,
		"total_dokter":      totalDokter,
		"total_staff":       totalStaff,
		"poli":              Poli,
	})
}
