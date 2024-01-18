// authentication/controller.go

// Registrasi <= /api/v1/resource/admin/registration

// di halaman ini terdapat tanggal lahir, yang menerima data berupa string (dd-mm-yyyy)

// di halaman ini dilakukan generate token, apabila belum ada siapapun yang masuk,
// maka akan digunakan token default yaitu (tokenAdmin)
// setelahnya akan dilakukan auto generate token 4 digit

package authentication

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputRegistrasi struct {
	Nama         string `json:"nama"`
	TanggalLahir string `json:"tanggal_lahir"`
	Gender       bool   `json:"gender"`
	Alamat       string `json:"alamat"`
	Password     string `json:"password"`
	Token        string `json:"token"`
}

// ================================================================

// Registration <= POST api/v1/resources/admin/registration
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
		Gender:       Registrasi.Gender,
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

	c.JSON(http.StatusOK, gin.H{
		"id":   User.ID,
		"nip":  User.Nip,
		"role": User.Role,
	})
}

func Nip(registrasi InputRegistrasi) (nip string) {
	// Kodifikasi NIP
	// 01 for admin
	// 01<bulan masuk><tahun masuk><bulan lahir><tahun lahir><2 digit angka random>\

	date := registrasi.TanggalLahir
	layout := "02-01-2006"
	parsedTanggalLahir, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	var bulanMasuk int
	var bulanLahir int

	tahunLahir := parsedTanggalLahir.Year()
	tahunLahirStr := strconv.Itoa(tahunLahir)
	bulanLahir = int(parsedTanggalLahir.Month())

	randomNumber := rand.Intn(100)
	masuk := time.Now()
	tahunMasuk := masuk.Year()
	tahunMasukStr := strconv.Itoa(tahunMasuk)
	bulanMasuk = int(masuk.Month())

	// Generate Nip
	nip = fmt.Sprintf("01%02d%2s%02d%2s%02d", bulanMasuk, tahunMasukStr[2:], bulanLahir, tahunLahirStr[2:], randomNumber)

	return nip
}
