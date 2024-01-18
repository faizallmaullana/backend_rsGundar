// GET SearchNik <= /api/v1/resources/pasien/:nik

package medical_record

import (
	"net/http"

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

	c.JSON(http.StatusFound, gin.H{
		"message": "Nik ditemukan",
		"status":  true,
	})
}
