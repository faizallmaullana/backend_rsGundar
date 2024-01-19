package base_on_page

import (
	"net/http"

	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

// get /pasienList/:idPasien
func RiwayatKunjungan(c *gin.Context) {
	var pasien models.Pasien
	if err := models.DB.Where("id = ?", c.Param("idPasien")).First(&pasien).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pasien tidak ditemukan"})
		return
	}

	var medical_record []models.MedicalRecord
	if err := models.DB.Where("id_pasien = ?", c.Param("idPasien")).Preload("dokter").Preload("pasien").Preload("dokter.poli").Preload("diagnosis").Find(&medical_record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data medicalRecord tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"nama_pasien":       pasien.Nama,
		"dataMedicalRecord": medical_record,
	})
}
