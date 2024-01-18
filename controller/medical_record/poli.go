// authentication/controller.go

// POST Registrasi <= /api/v1/resources/poli

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

// POST Registrasi <= /api/v1/resources/poli
func AddPoli(c *gin.Context) {
	var input InputPoli
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek jika ada nama poli yabg sama
	var dbPoli models.Poli
	if err := models.DB.Where("poli = ?", encryption.Encrypt(input.Poli)).First(&dbPoli).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Poli sudah terdaftar"})
		return
	}

	// generate id
	id := uuid.New().String()

	// enkripsi
	poliEnc := encryption.Encrypt(strings.ToLower(input.Poli))

	// create ke database
	poli := models.Poli{
		ID:   id,
		Poli: poliEnc,
	}

	models.DB.Create(&poli)
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil direkam"})
}
