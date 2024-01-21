// GET ListAllDokter <= /dokter/list/all

package medical_record

import (
	"net/http"
	"strings"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

func ListAllDokter(c *gin.Context) {
	var dokterList []models.Users
	if err := models.DB.Where("role = ?", encryption.Encrypt("dokter")).Preload("Profile").Find(&dokterList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database cant access"})
		return
	}

	var decrytedDokterList []map[string]interface{}
	for _, dokter := range dokterList {
		decryptedName := strings.Title(encryption.Decrypt(dokter.Profile.Nama))
		decrytedDokterList = append(decrytedDokterList, map[string]interface{}{
			"nama": decryptedName,
			"id":   dokter.ID,
		})
	}

	c.JSON(http.StatusOK, gin.H{"list": decrytedDokterList})
}
