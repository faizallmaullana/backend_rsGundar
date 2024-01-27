// halaman profile untuk profile admin dan staff pendaftaran

// GET Profile <= /api/v1/resources/profile/:user_id

package profile

import (
	"net/http"
	"strings"
	"time"

	"github.com/faizallmaullana/be_rsGundar/api/encryption"
	"github.com/faizallmaullana/be_rsGundar/api/models"
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

	// formated tanggal lahir
	// Assuming you have a time.Time variable named "myDate"
	myDate := user.Profile.TanggalLahir

	// Format the date to "02 01 06" layout
	formattedDate := myDate.Format("02-01-2006")

	// jika role user bukan dokter
	if role != "dokter" {
		c.JSON(http.StatusOK, gin.H{
			"id_user":       user.ID,
			"id_profile":    user.IDProfile,
			"nip":           user.Nip,
			"tanggal_lahir": formattedDate,
			"nama":          nama,
			"gender":        gender,
			"alamat":        alamat,
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
		"tanggal_lahir":     formattedDate,
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

	var userDB models.Users
	if err := models.DB.Where("id = ?", c.Param("user_id")).First(&userDB).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	var profileDB models.Profile
	if err := models.DB.Where("id = ?", userDB.IDProfile).First(&profileDB).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile tidak ditemukan"})
		return
	}

	// enkripsi
	nama := encryption.Encrypt(input.Nama)
	alamat := encryption.Encrypt(input.Alamat)

	// Load the UTC+7 (Indochina Time) location
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return
	}

	// Parse the date in the specified location
	var date string
	layout := "02-01-2006"
	if date != "" {
		date = input.TanggalLahir
	}

	parsedTanggalLahir, err := time.ParseInLocation(layout, date, location)
	if date != "" {
		if err != nil {
			return
		}
	}

	// convert request gender to bool
	var gender bool
	if input.Gender == "pria" {
		gender = true
	} else if input.Gender == "wanita" {
		gender = false
	}

	password, _ := encryption.HashPassword(input.Password)

	dataProfile := models.Profile{
		Nama:         nama,
		Alamat:       alamat,
		Gender:       gender,
		TanggalLahir: parsedTanggalLahir,
	}

	dataUser := models.Users{
		Password: password,
	}

	models.DB.Model(&profileDB).Update(&dataProfile)
	models.DB.Model(&userDB).Update(&dataUser)

	c.JSON(http.StatusCreated, gin.H{
		"user_id": userDB.ID,
		"nama":    encryption.Decrypt(profileDB.Nama),
		"gender":  profileDB.Gender,
	})
}
