// this is the medical_records models
// pasien, poli, diagnosis, income, temp_pendaftaran, medical_record

package models

import "time"

type Pasien struct {
	ID           string    `json:"id"`
	Nik          string    `json:"nik"`
	Nama         string    `json:"nama"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Gender       bool      `json:"gender"`
	Alamat       string    `json:"alamat"`
}

type Poli struct {
	ID        string    `json:"id"`
	Poli      string    `json:"poli"`
	CreatedAt time.Time `json:"created_at"`
	Total     int       `json:"total"`
}

type Diagnosis struct {
	ID        string `json:"id"`
	Diagnosis string `json:"diagnosis"`
	Total     int64  `json:"total"`
}

type TempPendaftaran struct {
	ID string `json:"id"`

	// foreign keys
	IDPasien string `json:"id_pasien"`
	IDDokter string `json:"id_dokter"`
	Biaya    string `json:"biaya"`

	// reference to
	Pasien Pasien        `json:"pisien" gorm:"primaryKey:IDPasien"`
	Dokter ProfileDokter `json:"profile_dokter" gorm:"primaryKey:IDokter"`
}

type MedicalRecord struct {
	ID        string    `json:"id"`
	Gejala    string    `json:"gejala"`
	Obat      string    `json:"obat"`
	Biaya     int       `json:"biaya"`
	CreatedAt time.Time `json:"created_at"`

	// foreign keys
	IDPasien    string `json:"id_pasien"`
	IDDokter    string `json:"id_dokter"`
	IDDiagnosis string `json:"id_diagnosis"`

	// reference to
	Pasien    Pasien        `json:"pasien" gorm:"primaryKey:IDPasien"`
	Dokter    ProfileDokter `json:"profile_dokter" gorm:"primaryKey:IDDokter"`
	Diagnosis Diagnosis     `json:"diagnosis" gorm:"primaryKey:IDDiagnosis"`
}

type Income struct {
	ID     string `json:"id"`
	Income int    `json:"income"`
}
