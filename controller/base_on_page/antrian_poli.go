package base_on_page

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

// get /antrianPoli/:idDokter
func AntrianPoli(c *gin.Context) {
	var Antrian []models.TempPendaftaran
	if err := models.DB.Where("id_dokter = ?", c.Param("idDokter")).Find(&Antrian).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data antrian tidak ditemukan"})
		return
	}

	// Dapatkan waktu mulai hari ini
	todayStart := time.Now().Truncate(24 * time.Hour)

	// Dapatkan waktu akhir hari ini
	todayEnd := todayStart.Add(24 * time.Hour).Add(-time.Second)

	var AntrianSelesai []models.MedicalRecord
	if err := models.DB.Where("created_at BETWEEN ? AND ?", todayStart, todayEnd).Find(&AntrianSelesai).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data medis tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"dalam_antrian":   Antrian,
		"selesai_antrian": AntrianSelesai,
	})
}
