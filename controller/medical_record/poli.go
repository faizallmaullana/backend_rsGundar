// authentication/controller.go

// GET GetAllPoli <= /api/v1/resources/poli
// POST AddPoli <= /api/v1/resources/poli

// di halaman ini terdapat tanggal lahir, yang menerima data berupa string (dd-mm-yyyy)

package medical_record

import (
	"net/http"
	"strings"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputPoli struct {
	Poli string `json:"poli"`
}

// GET GetAllPoli <= /api/v1/resources/poli
func GetAllPoli(c *gin.Context) {
	var poliList []models.Poli
	if err := models.DB.Find(&poliList).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// dekripsi
	var decryptedPoliList []map[string]interface{}
	for _, poli := range poliList {
		decryptedName := strings.Title(encryption.Decrypt(poli.Poli))
		decryptedPoliList = append(decryptedPoliList, map[string]interface{}{
			"name": decryptedName,
			"id":   poli.ID, // Assuming ID is the field name in the Poli struct
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"poli": decryptedPoliList,
	})
}

// POST AddPoli <= /api/v1/resources/poli
func AddPoli(c *gin.Context) {
	var input InputPoli
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputPoli := strings.ToLower(input.Poli)

	// cek jika ada nama poli yabg sama
	var dbPoli models.Poli
	if err := models.DB.Where("poli = ?", encryption.Encrypt(inputPoli)).First(&dbPoli).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Poli sudah terdaftar"})
		return
	}

	// generate id
	id := uuid.New().String()

	// enkripsi
	poliEnc := encryption.Encrypt(inputPoli)

	// create ke database
	poli := models.Poli{
		ID:   id,
		Poli: poliEnc,
	}

	models.DB.Create(&poli)
	c.JSON(http.StatusCreated, gin.H{"message": "data berhasil direkam"})
}
