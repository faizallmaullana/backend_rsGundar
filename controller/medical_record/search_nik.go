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
		c.JSON(http.StatusOK, gin.H{
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

	// medicalRecord
	var MedicalRecords []models.MedicalRecord
	db := models.DB.Where("id_pasien = ?", Pasien.ID)
	dbPreload := db.Preload("Diagnosis").Preload("Dokter.Profile").Preload("Dokter.ProfileDokter.Poli")
	if err := dbPreload.Find(&MedicalRecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data medis tidak ditemukan"})
		return
	}

	var decryptedData []map[string]interface{}
	for _, record := range MedicalRecords {
		dokter := strings.Title(encryption.Decrypt(record.Dokter.Profile.Nama))
		poli := strings.Title(encryption.Decrypt(record.Dokter.ProfileDokter.Poli.Poli))
		diagnosis := strings.Title(encryption.Decrypt(record.Diagnosis.Diagnosis))
		decryptedData = append(decryptedData, map[string]interface{}{
			"id":        record.ID,
			"dokter":    dokter,
			"poli":      poli,
			"diagnosis": diagnosis,
		})
	}

	// formated tanggal lahir
	// Assuming you have a time.Time variable named "myDate"
	myDate := Pasien.TanggalLahir

	// Format the date to "02 01 06" layout
	formattedDate := myDate.Format("02-01-2006")

	c.JSON(http.StatusOK, gin.H{
		"message":       "Nik ditemukan",
		"status":        true,
		"id":            Pasien.ID,
		"nik":           Pasien.Nik,
		"alamat":        alamat,
		"nama":          nama,
		"gender":        gender,
		"data_medis":    decryptedData,
		"tanggal_lahir": formattedDate,
	})
}

// ======================== hapuseun
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
