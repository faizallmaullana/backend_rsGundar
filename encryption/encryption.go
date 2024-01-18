package encryption

// Fungsi untuk mengenkripsi teks dengan menggunakan nilai geser dari kunci
func Encrypt(text string) string {
	key := "hdafldhfadfkjbasdofaosiflksdnflk"
	result := ""
	keyLen := len(key)

	for i := 0; i < len(text); i++ {
		char := text[i]
		shift := key[i%keyLen]

		// Periksa apakah karakter adalah huruf (A-Z atau a-z)
		if char >= 'A' && char <= 'Z' {
			char = (char-'A'+byte(shift-'A'))%26 + 'A'
		} else if char >= 'a' && char <= 'z' {
			char = (char-'a'+byte(shift-'a'))%26 + 'a'
		}
		result += string(char)
	}
	return result
}

// Fungsi untuk mendekripsi teks yang telah dienkripsi dengan menggunakan nilai geser dari kunci
func Decrypt(text string) string {
	key := "hdafldhfadfkjbasdofaosiflksdnflk"

	result := ""
	keyLen := len(key)

	for i := 0; i < len(text); i++ {
		char := text[i]
		shift := key[i%keyLen]

		// Periksa apakah karakter adalah huruf (A-Z atau a-z)
		if char >= 'A' && char <= 'Z' {
			char = (char-'A'+26-byte(shift-'A'))%26 + 'A'
		} else if char >= 'a' && char <= 'z' {
			char = (char-'a'+26-byte(shift-'a'))%26 + 'a'
		}
		result += string(char)
	}
	return result
}
