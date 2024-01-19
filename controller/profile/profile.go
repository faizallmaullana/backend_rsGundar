// halaman profile untuk profile admin dan staff pendaftaran

// GET Profile <= /api/v1/resources/profile/:user_id

package profile

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/faizallmaullana/be_rsGundar/encryption"
	"github.com/faizallmaullana/be_rsGundar/models"
	"github.com/gin-gonic/gin"
)

type InputProfile struct {
	Nama         string `json:"nama"`
	Gender       string `json:"gender"`
	Alamat       string `json:"alamat"`
	TanggalLahir string `json:"tanggal_lahir"`
	Password     string `json:"password"`
}

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

// PUT ubah profile id "pakai parameter"
func UpdateProfile(c *gin.Context) {
	var input InputProfile
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var profileDB models.Profile
	if err := models.DB.Where("id = ?", c.Param("id_profile")).First(&profileDB); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile tidak ditemukan"})
		return
	}

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

	dataProfile := models.Profile{
		Nama:         nama,
		Alamat:       alamat,
		Gender:       gender,
		TanggalLahir: parsedTanggalLahir,
	}

	models.DB.Model(&profileDB).Update(&dataProfile)

	c.JSON(http.StatusCreated, gin.H{
		"id_profile": dataProfile.ID,
		"nama":       encryption.Decrypt(dataProfile.Nama),
	})
}
