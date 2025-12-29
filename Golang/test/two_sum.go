package test

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // Buat map: key: angka, value: indeks
	for i, num := range nums { // Loop dari awal sampai akhir
		complement := target - num // Hitung angka yang dibutuhkan (complement)
		if index, found := numMap[complement]; found { // Cek di map, ada gak complement-nya?
			return []int{index, i} // Kalo ada, BALIKIN INDEKSNYA (index adalah indeks si complement, i adalah indeks si num saat ini)
		}
		numMap[num] = i // Kalo gak ada, simpen num dan indeksnya ke map
	}
	return []int{} // Harusnya ketemu, tapi kalo gak ketemu, balikin array kosong (sesuai soal)
}