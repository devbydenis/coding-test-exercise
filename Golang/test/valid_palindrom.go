package test

func ValidPalindrome(s string) bool { // kasur rusak
	left := 0 // pointer kiri
	right := len(s) - 1 // pointer kanan

	for left < right { // pastiin left selalu lebih kecil dari right biar ga cross gitu nanti
		for left < right && !isAlpanumeric(s[left]) { 
			left ++ // kita longkap kalo ada karakter yang bukan alfanumerik
		}

		for left < right && !isAlpanumeric(s[right]) {
			right --
		}

		if toLower(s[left]) != toLower(s[right]) { // kita bikin tiap karakter lowercase dulu baru kita bandingin
			return false
		}

		left++ // geser dari kiri kekanan
		right-- // geser dari kanan kekiri
	}

	return true
}

func isAlpanumeric(c byte) bool { // tiap karakter itu adalah byte kalo di go
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32 	// biar dia lowercase
	}

	return c
}


