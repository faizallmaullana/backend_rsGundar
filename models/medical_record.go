// this is the medical_records models
// pasien, poli, diagnosis, income, temp_pendaftaran, medical_record

package models

import "time"

type Pasien struct {
	ID           string    `json:"id" gorm:"primary_key"`
	Nik          string    `json:"nik"`
	Nama         string    `json:"nama"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Gender       bool      `json:"gender"`
	Alamat       string    `json:"alamat"`
}

type Poli struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Poli      string    `json:"poli"`
	CreatedAt time.Time `json:"created_at"`
	Total     int       `json:"total"`
}

type Diagnosis struct {
	ID        string `json:"id" gorm:"primary_key"`
	Diagnosis string `json:"diagnosis"`
	Total     int    `json:"total"`
}

type TempPendaftaran struct {
	ID string `json:"id"`

	// foreign keys
	IDPasien string `json:"id_pasien" gorm:"primary_key"`
	IDDokter string `json:"id_dokter"`
	Biaya    int    `json:"biaya"`

	// reference to
	Pasien Pasien `json:"pasien" gorm:"foreignKey:IDPasien"`
	Dokter Users  `json:"dokter" gorm:"foreignKey:IDDokter"`
}

type MedicalRecord struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Gejala    string    `json:"gejala"`
	Obat      string    `json:"obat"`
	Biaya     int       `json:"biaya"`
	CreatedAt time.Time `json:"created_at"`

	// foreign keys
	IDPasien    string `json:"id_pasien"`
	IDDokter    string `json:"id_dokter"`
	IDDiagnosis string `json:"id_diagnosis"`

	// reference to
	Pasien    Pasien        `json:"pasien" gorm:"foreignKey:IDPasien"`
	Dokter    ProfileDokter `json:"profile_dokter" gorm:"foreignKey:IDDokter"`
	Diagnosis Diagnosis     `json:"diagnosis" gorm:"foreignKey:IDDiagnosis"`
}

type Income struct {
	ID     string `json:"id" gorm:"primary_key"`
	Income int    `json:"income"`
}
