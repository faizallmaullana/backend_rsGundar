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
