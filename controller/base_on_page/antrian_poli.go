package base_on_page

import (
	"net/http"
	"strings"
	"time"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

// get /antrianPoli/:idDokter
func AntrianPoli(c *gin.Context) {
	var Antrian []models.TempPendaftaran
	dbAntrian := models.DB.Where("id_dokter = ?", c.Param("idDokter"))
	dbPreAntrian := dbAntrian.Preload("Pasien")
	if err := dbPreAntrian.Find(&Antrian).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data antrian tidak ditemukan"})
		return
	}

	var decryptedAntrian []map[string]interface{}
	for _, antrian := range Antrian {
		decryptedName := strings.Title(encryption.Decrypt(antrian.Pasien.Nama))
		decryptedAntrian = append(decryptedAntrian, map[string]interface{}{
			"nama": decryptedName,
			"id":   antrian.ID,
		})
	}
	// Dapatkan waktu mulai hari ini
	todayStart := time.Now().Truncate(24 * time.Hour)

	// Dapatkan waktu akhir hari ini
	todayEnd := todayStart.Add(24 * time.Hour).Add(-time.Second)

	var AntrianSelesai []models.MedicalRecord
	dbAntrianSelesai := models.DB.Where("created_at BETWEEN ? AND ?", todayStart, todayEnd)
	if err := dbAntrianSelesai.Find(&AntrianSelesai).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data medis tidak ditemukan"})
		return
	}

	var kunjunganSelesai []map[string]interface{}
	for _, antrian := range AntrianSelesai {
		decryptedName := strings.Title(encryption.Decrypt(antrian.Pasien.Nama))
		kunjunganSelesai = append(kunjunganSelesai, map[string]interface{}{
			"name": decryptedName,
			"id":   antrian.ID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"selesai_diperiksa":       AntrianSelesai,
		"belum_selesai_diperiksa": decryptedAntrian,
	})
}
