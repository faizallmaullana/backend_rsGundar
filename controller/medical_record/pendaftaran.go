// POST PendaftaranMedicalRecord <= /medicalRecord/pendaftaran/:pasien_id
// GET DataFromPendaftaran <= /data/from/pendaftaran/id_pendaftaran

package medical_record

import (
	"net/http"
	"strings"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputPendaftaran struct {
	DokterID string `json:"dokter_id"`
	Biaya    int    `json:"biaya"`
}

// POST PendaftaranMedicalRecord <= /medicalRecord/pendaftaran/:pasien_id
func PendafataranMedicalRecord(c *gin.Context) {
	var input InputPendaftaran
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Pasien models.Pasien
	if err := models.DB.Where("id = ?", c.Param("pasien_id")).First(&Pasien).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pasien tidak ditemukan"})
		return
	}

	var profileDokter models.Users
	if err := models.DB.Where("id = ?", input.DokterID).Preload("Profile").Preload("ProfileDokter.Poli").First(&profileDokter).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dokter tidak ditemukan"})
		return
	}

	// generate id
	id := uuid.New().String()

	dataPendaftaran := models.TempPendaftaran{
		ID:       id,
		IDDokter: input.DokterID,
		IDPasien: Pasien.ID,
		Biaya:    input.Biaya,
	}

	models.DB.Create(&dataPendaftaran)

	c.JSON(http.StatusCreated, gin.H{
		"message":        "sukses",
		"id_pendaftaran": dataPendaftaran.ID,
		"id_pasien":      dataPendaftaran.IDPasien,
		"id_dokter":      dataPendaftaran.IDDokter,
		"id_poli":        profileDokter.ProfileDokter.IDPoli,
	})
}

// GET DataFromPendaftaran <= /data/from/pendaftaran/:id_pendaftaran
func DataFromPendaftaran(c *gin.Context) {
	var TempPendaftaran models.TempPendaftaran
	if err := models.DB.Where("id = ?", c.Param("id_pendaftaran")).Preload("Pasien").First(&TempPendaftaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "tidak bisa mengakses database"})
		return
	}

	// formated tanggal lahir
	// Assuming you have a time.Time variable named "myDate"
	myDate := TempPendaftaran.Pasien.TanggalLahir

	// Format the date to "02 01 06" layout
	formattedDate := myDate.Format("02-01-2006")

	// dektripsi
	nama := strings.Title(encryption.Decrypt(TempPendaftaran.Pasien.Nama))

	c.JSON(http.StatusOK, gin.H{
		"id":            TempPendaftaran.ID,
		"id_dokter":     TempPendaftaran.IDDokter,
		"id_pasien":     TempPendaftaran.IDPasien,
		"biaya":         TempPendaftaran.Biaya,
		"nama":          nama,
		"tanggal_lahir": formattedDate,
	})
}
