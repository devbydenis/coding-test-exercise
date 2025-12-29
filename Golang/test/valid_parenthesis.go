package test

func IsValid(s string) bool {
	match := map[rune]rune{ // Map untuk cocokin kurung tutup -> kurung buka
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{} // Slice di Go berfungsi sebagai Stack (push/pop)

	
	for _, char := range s { // Loop setiap karakter (rune)
		if openBracket, ok := match[char]; ok { // Cek: Apakah char adalah Kurung TUTUP?
			if len(stack) == 0 || stack[len(stack)-1] != openBracket { // 1.) Cek Stack kosong? (Ada kurung tutup tanpa kurung buka) 2.) Cek elemen teratas Stack cocok gak sama pasangannya?
				return false
			}
	
			stack = stack[:len(stack)-1] // Kalau cocok, Pop/Keluarkan elemen teratas dari Stack
		} else {
			stack = append(stack, char) // Kalau char adalah Kurung BUKA, Push/Masukkan ke Stack
		}
	}

	return len(stack) == 0 // Valid jika dan hanya jika setelah selesai loop, Stack KOSONG
}
