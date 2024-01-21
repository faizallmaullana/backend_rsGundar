// POST PendaftaranMedicalRecord <= /medicalRecord/pendaftaran

package medical_record

import (
	"net/http"

	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputPendaftaran struct {
	DokterID string `json:"dokter_id"`
	PasienID string `json:"pasien_id"`
	Biaya    string `json:"biaya"`
}

// POST PendaftaranMedicalRecord <= /medicalRecord/pendaftaran
func PendafataranMedicalRecord(c *gin.Context) {
	var input InputPendaftaran
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var profileDokter models.Users
	if err := models.DB.Where("id = ?", input.DokterID).Preload("Dokter").First(&profileDokter).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dokter tidak ditemukan"})
		return
	}

	// generate id
	id := uuid.New().String()

	dataPendaftaran := models.TempPendaftaran{
		ID:       id,
		IDDokter: input.DokterID,
		IDPasien: input.PasienID,
		Biaya:    input.Biaya,
	}

	models.DB.Create(&dataPendaftaran)

	c.JSON(http.StatusCreated, gin.H{
		"id_pendaftaran": dataPendaftaran.ID,
		"id_pasien":      dataPendaftaran.IDPasien,
		"id_dokter":      dataPendaftaran.IDDokter,
		"id_poli":        profileDokter.ProfileDokter.IDPoli,
	})
}
