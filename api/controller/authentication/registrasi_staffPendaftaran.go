// authentication/controller.go

// POST Registrasi <= /api/v1/resource/staffPendaftaran/registration

// di halaman ini terdapat tanggal lahir, yang menerima data berupa string (dd-mm-yyyy)

// di halaman ini dilakukan generate token, apabila belum ada siapapun yang masuk,
// maka akan digunakan token default yaitu (tokenAdmin)
// setelahnya akan dilakukan auto generate token 4 digit

// NOTE: pada tambah dokter belum ditambahkan id poli

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

// POST Registration <= POST api/v1/resources/staffPendaftaran/registration
func RegistrasiStaffPendaftaran(c *gin.Context) {

	// request handler
	var Registrasi InputRegistrasi
	if err := c.ShouldBindJSON(&Registrasi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generate the id
	idUser := uuid.New().String()
	idProfile := uuid.New().String()

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
	role := encryption.Encrypt("staffPendaftaran")
	nama := encryption.Encrypt(strings.ToLower(Registrasi.Nama))
	alamat := encryption.Encrypt(strings.ToLower(Registrasi.Alamat))
	password, _ := encryption.HashPassword("default")

	// convert request gender to bool
	var gender bool
	if Registrasi.Gender == "pria" {
		gender = true
	} else if Registrasi.Gender == "wanita" {
		gender = false
	}

	// save data for users table
	User := models.Users{
		ID:        idUser,
		Nip:       nip,
		Password:  password,
		Role:      role,
		IDProfile: idProfile,
	}

	// save data for profile table
	Profile := models.Profile{
		ID:           idProfile,
		Gender:       gender,
		Nama:         nama,
		Alamat:       alamat,
		TanggalLahir: parsedTanggalLahir,
	}

	// save data to the database
	models.DB.Create(&User)
	models.DB.Create(&Profile)

	c.JSON(http.StatusCreated, gin.H{
		"id":   User.ID,
		"nip":  User.Nip,
		"role": encryption.Decrypt(User.Role),
	})
}
