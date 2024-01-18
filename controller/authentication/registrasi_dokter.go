// authentication/controller.go

// POST Registrasi <= /api/v1/resource/dokter/registration

// di halaman ini terdapat tanggal lahir, yang menerima data berupa string (dd-mm-yyyy)

package authentication

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ================================================================

// Registration <= POST api/v1/resources/dokter/registration
func RegistrasiDokter(c *gin.Context) {

	// request handler
	var Registrasi InputRegistrasi
	if err := c.ShouldBindJSON(&Registrasi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generate the id
	idUser := uuid.New().String()
	idProfile := uuid.New().String()
	idProfileDokter := uuid.New().String()

	nip := Nip(Registrasi)

	date := Registrasi.TanggalLahir
	layout := "02-01-2006"

	// Load the UTC+7 (Indochina Time) location
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Parse the date in the specified location
	parsedTanggalLahir, err := time.ParseInLocation(layout, date, location)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	// enkripsi
	role := encryption.Encrypt("dokter")
	nama := encryption.Encrypt(strings.ToLower(Registrasi.Nama))
	alamat := encryption.Encrypt(strings.ToLower(Registrasi.Alamat))
	spesialisasi := encryption.Encrypt(strings.ToLower(Registrasi.Spesialisasi))
	password, _ := encryption.HashPassword(Registrasi.Password)

	// save data for users table
	User := models.Users{
		ID:              idUser,
		Nip:             nip,
		Password:        password,
		Role:            role,
		IDProfile:       idProfile,
		IDProfileDokter: idProfileDokter,
	}

	// save data for profile table
	Profile := models.Profile{
		ID:           idProfile,
		Gender:       Registrasi.Gender,
		Nama:         nama,
		Alamat:       alamat,
		TanggalLahir: parsedTanggalLahir,
	}

	// save data for profile dokter table
	ProfileDokter := models.ProfileDokter{
		ID:           idProfileDokter,
		Spesialisasi: spesialisasi,
		IDPoli:       Registrasi.PoliID,
	}

	// save data to the database
	models.DB.Create(&User)
	models.DB.Create(&Profile)
	models.DB.Create(&ProfileDokter)

	c.JSON(http.StatusOK, gin.H{
		"id":   User.ID,
		"nip":  User.Nip,
		"role": encryption.Decrypt(User.Role),
	})
}
