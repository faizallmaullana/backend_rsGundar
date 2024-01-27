// POST MedicalRecord <= /dokter/medicalRecord/:idTempPendaftaran

package base_on_page

import (
	"net/http"
	"strings"

	"github.com/faizallmaullana/be_rsGundar/api/encryption"
	"github.com/faizallmaullana/be_rsGundar/api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputMedicalRecord struct {
	Gejala    string `json:"gejala"`
	Diagnosis string `json:"diagnosis"`
	Obat      string `json:"obat"`
}

// PUSH MedicalRecord <= /dokter/medicalRecord/:idTempPendaftaran
func MedicalRecord(c *gin.Context) {
	var input InputMedicalRecord
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var TempPendaftaran models.TempPendaftaran
	dbTempPendaftaran := models.DB.Where("id = ?", c.Param("idTempPendaftaran"))
	if err := dbTempPendaftaran.Preload("Dokter.ProfileDokter").First(&TempPendaftaran).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data pendaftaran tidak ditemukan"})
		return
	}

	// generate id
	id := uuid.New().String()

	strDiagnosis := strings.ToLower(input.Diagnosis)

	// enkripsi
	gejala := encryption.Encrypt(input.Gejala)
	diagnosis := encryption.Encrypt(strDiagnosis)
	obat := encryption.Encrypt(input.Obat)

	var dbDiagnosis models.Diagnosis
	var DiagnosisID string
	if err := models.DB.Where("diagnosis = ?", diagnosis).First(&dbDiagnosis).Error; err == nil {
		// If a record is found, use the existing record's ID
		DiagnosisID = dbDiagnosis.ID
	} else {
		// If no record found, create a new one
		idDiagnosis := uuid.New().String()
		dataDiagnosis := models.Diagnosis{
			ID:        idDiagnosis,
			Diagnosis: diagnosis,
		}

		models.DB.Create(&dataDiagnosis)
		DiagnosisID = dataDiagnosis.ID
	}

	var Income models.Income
	if err := models.DB.First(&Income).Error; err != nil {
		idIncome := uuid.New().String()
		dataIncome := models.Income{
			ID:     idIncome,
			Income: TempPendaftaran.Biaya,
		}

		models.DB.Create(&dataIncome)
	} else {
		newIncomeValue := Income.Income + TempPendaftaran.Biaya
		models.DB.Model(&Income).Update("income", newIncomeValue)
	}

	var poli models.Poli
	dbPoli := models.DB.Where("id = ?", TempPendaftaran.Dokter.ProfileDokter.IDPoli)
	if err := dbPoli.First(&poli).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "data poli tidak ditemukan"})
		return
	}

	dataMedicalRecord := models.MedicalRecord{
		ID:          id,
		Gejala:      gejala,
		Obat:        obat,
		IDPasien:    TempPendaftaran.IDPasien,
		IDDokter:    TempPendaftaran.IDDokter,
		IDDiagnosis: DiagnosisID,
		Biaya:       TempPendaftaran.Biaya,
	}

	poli.Total++
	models.DB.Save(poli)

	models.DB.Create(&dataMedicalRecord)
	models.DB.Where("id = ?", TempPendaftaran.ID).Delete(&TempPendaftaran)

	c.JSON(http.StatusOK, gin.H{
		"message":           "sukses",
		"id_medical_record": dataMedicalRecord.ID,
	})
}
