// here is model for the users of the app
// users, profile, profile_dokter, db_token

package models

import "time"

type Users struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Nip       string    `json:"nip"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`

	// foreign keys
	IDProfile       string `json:"id_profile"`
	IDProfileDokter string `json:"id_profile_dokter"`

	// references to
	Profile       Profile       `json:"profile" gorm:"foreignKey:IDProfile"`
	ProfileDokter ProfileDokter `json:"profile_dokter" gorm:"foreignKey:IDProfileDokter"`
}

type Profile struct {
	ID           string    `json:"id" gorm:"primary_key"`
	Nama         string    `json:"nama"`
	Gender       bool      `json:"gender"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Alamat       string    `json:"alamat"`
}

type ProfileDokter struct {
	ID           string `json:"id" gorm:"primary_key"`
	Spesialisasi string `json:"spesialisasi"`

	// foreign keys
	IDPoli string `json:"id_poli"`

	// references to
	Poli Poli `json:"poli" gorm:"foreignKey:IDPoli"`
}

type DBToken struct {
	ID    string `json:"id" gorm:"primary_key"`
	Token string `json:"token"`
}
