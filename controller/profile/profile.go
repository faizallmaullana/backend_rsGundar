// halaman profile untuk profile admin dan staff pendaftaran

// GET Profile <= /api/v1/resources/profile/:user_id

package profile

import (
	"net/http"
	"strings"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

// GET Profile <= /api/v1/resources/profile/:user_id
func Profile(c *gin.Context) {
	var user models.Users
	if err := models.DB.Preload("Profile").Preload("ProfileDokter").Preload("ProfileDokter.Poli").Where("id = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// dekripsi
	role := encryption.Decrypt(user.Role)
	nama := strings.Title(encryption.Decrypt(user.Profile.Nama))
	alamat := strings.Title(encryption.Decrypt(user.Profile.Alamat))
	spesialisasi := strings.Title(encryption.Decrypt(user.ProfileDokter.Spesialisasi))
	poli := strings.Title(encryption.Decrypt(user.ProfileDokter.Poli.Poli))

	var gender string
	if !user.Profile.Gender {
		gender = "Wanita"
	} else {
		gender = "Pria"
	}

	// jika role user bukan dokter
	if role != "dokter" {
		c.JSON(http.StatusOK, gin.H{
			"id_user":    user.ID,
			"id_profile": user.IDProfile,
			"nip":        user.Nip,
			"nama":       nama,
			"gender":     gender,
			"alamat":     alamat,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_user":           user.ID,
		"id_profile":        user.IDProfile,
		"id_profile_dokter": user.IDProfileDokter,
		"nip":               user.Nip,
		"nama":              nama,
		"gender":            gender,
		"alamat":            alamat,
		"spesialisasi":      spesialisasi,
		"poli":              poli,
	})
}
