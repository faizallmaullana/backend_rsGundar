// authentication/login.go

// POST Login <= /api/v1/resources/login

// di halaman ini terdapat fungsi untuk melakukan login
// login dilakukan seperti biasa
// program memerlukan NIP dan Password untuk masuk

package authentication

import (
	"fmt"
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
	fmt.Println("test    1")
	var input InputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("test    2")

	// cek apakah nip terdaftar
	var user models.Users
	if err := models.DB.Where("nip = ? ", input.Nip).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "NIP tidak terdaftar"})
		return
	}

	fmt.Println("test    3")

	// dekripsi
	role := encryption.Decrypt(user.Role)

	fmt.Println("test    4")

	// cek password
	matchPassword := encryption.CheckPasswordHash(input.Password, user.Password)

	fmt.Println(matchPassword)
	fmt.Println("test    5")

	if !matchPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password salah"})
		return
	}
	fmt.Println("test    6")

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Berhasil",
		"id":      user.ID,
		"role":    role,
		"nip":     user.Nip,
	})
}
