// GET SearchNik <= /api/v1/resources/pasien/:nik
// GET DataPasienSatuan <= /pasien/satuan/:id_pasien

package medical_record

import (
	"net/http"
	"strings"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

func SearchNik(c *gin.Context) {
	var Pasien models.Pasien
	if err := models.DB.Where("nik = ?", c.Param("nik")).First(&Pasien).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Nik tidak ditemukan",
			"status":  false,
		})
		return
	}

	// dektripsi
	nama := strings.Title(encryption.Decrypt(Pasien.Nama))
	alamat := strings.Title(encryption.Decrypt(Pasien.Alamat))

	var gender string
	if !Pasien.Gender {
		gender = "Wanita"
	} else {
		gender = "Pria"
	}

	c.JSON(http.StatusFound, gin.H{
		"message": "Nik ditemukan",
		"status":  true,
		"id":      Pasien.ID,
		"nik":     Pasien.Nik,
		"alamat":  alamat,
		"nama":    nama,
		"gender":  gender,
	})
}

// data pasien satuan
// GET DataPasienSatuan <= /pasien/satuan/:id_pasien
func DataPasienSatuan(c *gin.Context) {
	var Pasien models.Pasien
	if err := models.DB.Where("id = ?", c.Param("id_pasien")).First(&Pasien).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data pasien tidak ditemukan"})
		return
	}

	var MedicalRecords []models.MedicalRecord
	if err := models.DB.Where("id_pasien = ?", Pasien.ID).Preload("Pasien").Preload("Dokter").Preload("Diagnosis").Find(&MedicalRecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data medis tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data_pasien": Pasien,
		"data_medis":  MedicalRecords,
	})
}
