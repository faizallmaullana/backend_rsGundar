// authentication/controller.go

// POST Registrasi <= /api/v1/resource/registration/admin

// di halaman ini terdapat tanggal lahir, yang menerima data berupa string (dd-mm-yyyy)

package authentication

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/faizallmaullana/be_rsGundar/api/encryption"
	"github.com/faizallmaullana/be_rsGundar/api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputRegistrasi struct {
	Nama         string `json:"nama"`
	TanggalLahir string `json:"tanggal_lahir"`
	Gender       string `json:"gender"`
	Alamat       string `json:"alamat"`
	Password     string `json:"password"`
	Token        string `json:"token"`
	Spesialisasi string `json:"spesialisasi"`
	PoliID       string `json:"poli_id"`
}

// ================================================================

// Registration <= POST api/v1/resources/registration/admin
func Registrasi(c *gin.Context) {

	// request handler
	var Registrasi InputRegistrasi
	if err := c.ShouldBindJSON(&Registrasi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// deklarasi token <= tokenAdmin jika belum ada siapapun yang daftar
	var token string
	var CekToken models.DBToken
	if err := models.DB.First(&CekToken).Error; err != nil {
		token = "tokenAdmin"
	} else {
		token = CekToken.Token
	}

	// cek token
	if Registrasi.Token != token {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TokenSalah"})
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
	role := encryption.Encrypt("admin")
	nama := encryption.Encrypt(strings.ToLower(Registrasi.Nama))
	alamat := encryption.Encrypt(strings.ToLower(Registrasi.Alamat))

	// convert request gender to bool
	var gender bool
	if Registrasi.Gender == "pria" {
		gender = true
	} else if Registrasi.Gender == "wanita" {
		gender = false
	}

	password, _ := encryption.HashPassword(Registrasi.Password)

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

	// generate tokenBaru
	if err := models.DB.First(&CekToken).Error; err != nil {
		idToken := uuid.New().String()

		generateToken := rand.Intn(9000) + 1000
		generatedToken := models.DBToken{
			ID:    idToken,
			Token: strconv.Itoa(generateToken),
		}

		models.DB.Create(generatedToken)
	} else {
		generateToken := rand.Intn(9000) + 1000
		generatedToken := models.DBToken{
			Token: strconv.Itoa(generateToken),
		}

		models.DB.Model(&CekToken).Update(generatedToken)
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":   User.ID,
		"nip":  User.Nip,
		"role": encryption.Decrypt(User.Role),
	})
}
