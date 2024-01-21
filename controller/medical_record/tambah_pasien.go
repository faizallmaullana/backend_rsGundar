// POST TambahPasien <= /pasien/tambah

package medical_record

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

type InputPasien struct {
	Nik          string `json:"nik"`
	Nama         string `json:"nama"`
	TanggalLahir string `json:"tanggal_lahir"`
	Gender       string `json:"gender"`
	Alamat       string `json:"alamat"`
}

func TambahPasien(c *gin.Context) {
	var input InputPasien
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var pasien models.Pasien
	if err := models.DB.Where("nik = ?", input.Nik).First(&pasien).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Nik sudah terdaftar"})
		return
	}

	// autoGenerate
	id := uuid.New().String()

	// enkripsi
	nama := encryption.Encrypt(input.Nama)
	alamat := encryption.Encrypt(input.Alamat)

	date := input.TanggalLahir
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

	// convert request gender to bool
	var gender bool
	if input.Gender == "pria" {
		gender = true
	} else if input.Gender == "wanita" {
		gender = false
	}

	dataPasien := models.Pasien{
		ID:           id,
		Nik:          input.Nik,
		Nama:         nama,
		Alamat:       alamat,
		TanggalLahir: parsedTanggalLahir,
		Gender:       gender,
	}

	models.DB.Create(&dataPasien)

	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"nik":  dataPasien.Nik,
		"nama": strings.Title(encryption.Decrypt(dataPasien.Nama)),
	})
}
