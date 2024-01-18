// authentication/login.go

// POST Login <= /api/v1/resources/login

// di halaman ini terdapat fungsi untuk melakukan login
// login dilakukan seperti biasa
// program memerlukan NIP dan Password untuk masuk

package authentication

import (
	"net/http"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

type InputLogin struct {
	Nip      string `json:"nip"`
	Password string `json:"password"`
}

// Login <= /api/v1/resources/login
func Login(c *gin.Context) {
	var input InputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek apakah nip terdaftar
	var user models.Users
	if err := models.DB.Where("nip = ? ", input.Nip).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "NIP tidak terdaftar"})
		return
	}

	// dekripsi
	role := encryption.Decrypt(user.Role)

	// cek password
	matchPassword := encryption.CheckPasswordHash(input.Password, user.Password)

	if !matchPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password salah"})
		return
	}

	// debuging========

	var profile models.Profile
	if err := models.DB.Where("id = ? ", user.IDProfile).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Profile tidak terdaftar"})
		return
	}

	// ===============

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Berhasil",
		"id":      user.ID,
		"role":    role,
	})
}