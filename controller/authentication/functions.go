package authentication

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

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
	nip = fmt.Sprintf("%02d%2s%02d%2s%02d", bulanMasuk, tahunMasukStr[2:], bulanLahir, tahunLahirStr[2:], randomNumber)

	return nip
}

func titleCase(data string) string {
	return strings.Title(data)
}
