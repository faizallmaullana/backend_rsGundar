package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/faizallmaullana/be_rsGundar/api/controller/authentication"
	"github.com/faizallmaullana/be_rsGundar/api/controller/base_on_page"
	"github.com/faizallmaullana/be_rsGundar/api/controller/medical_record"
	"github.com/faizallmaullana/be_rsGundar/api/controller/profile"
	"github.com/faizallmaullana/be_rsGundar/api/controller/statistik"
	"github.com/faizallmaullana/be_rsGundar/api/models"
)

// initilaize the cors middleware
var corsConfig = cors.DefaultConfig()

func init() {
	// allow all origins
	corsConfig.AllowAllOrigins = true
}

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	// connect to database
	models.ConnectDatabase()
	r.Use(cors.New(corsConfig))

	// ROUTES

	// authentication
	r.POST("/api/v1/resources/registration/dokter", authentication.RegistrasiDokter)                     // tested
	r.POST("/api/v1/resources/registration/admin", authentication.Registrasi)                            // tested
	r.POST("/api/v1/resources/registration/staffPendaftaran", authentication.RegistrasiStaffPendaftaran) // tested
	r.POST("/api/v1/resources/login", authentication.Login)
	r.GET("/api/v1/resources/dataPegawai", authentication.GetAllPegawai) // tested with

	// poli
	r.POST("/api/v1/resources/poli", medical_record.AddPoli)   // tested
	r.GET("/api/v1/resources/poli", medical_record.GetAllPoli) // tested

	// profile
	r.GET("/api/v1/resources/profile/:user_id", profile.Profile)        // tested
	r.POST("/api/v1/resources/profile/:user_id", profile.UpdateProfile) // tested

	// medical records
	r.GET("/api/v1/resources/medical_record/all", medical_record.GetAllMedicalRecord)                          //
	r.GET("/api/v1/resources/pasien/:nik", medical_record.SearchNik)                                           // tested
	r.GET("/api/v1/resources/pasienID/:id", medical_record.SearchNik)                                          // tested
	r.POST("/api/v1/resources/pasien/tambah", medical_record.TambahPasien)                                     // tested
	r.POST("/api/v1/resources/medicalRecord/pendaftaran/:pasien_id", medical_record.PendafataranMedicalRecord) // tested
	r.GET("/api/v1/resources/dokter/list/all", medical_record.ListAllDokter)                                   // tested
	r.GET("/api/v1/resources/data/from/pendaftaran/:id_pendaftaran", medical_record.DataFromPendaftaran)       // tested
	r.GET("/api/v1/resources/pasien/satuan/:id_pasien", medical_record.DataPasienSatuan)                       // tested
	r.POST("/api/v1/resources/dokter/medicalRecord/:idTempPendaftaran", base_on_page.MedicalRecord)
	r.GET("/api/v1/resources/dokter/pemeriksaan/list/:idDokter", medical_record.GetDokterBaseID) // tested

	r.GET("/api/v1/resources/antrianPoli/:idDokter", base_on_page.AntrianPoli)     // tested but need some
	r.GET("/api/v1/resources/pasienList/:idPasien", base_on_page.RiwayatKunjungan) // tested with

	r.GET("/api/v1/resources/token", base_on_page.GetToken) // tested

	// statistik
	r.GET("/api/v1/resources/statistik", statistik.Statistik)

	// run the server
	r.Run(":3200")
}
