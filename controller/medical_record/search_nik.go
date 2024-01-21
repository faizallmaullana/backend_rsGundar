// GET SearchNik <= /api/v1/resources/pasien/:nik

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
