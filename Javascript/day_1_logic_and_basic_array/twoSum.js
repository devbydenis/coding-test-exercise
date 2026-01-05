/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */

// Pake Map atau Object (Hash Map)
export var twoSum = function (nums, target) {
  const numMap = new Map(); // Buat map: (key: angka, value: indeks)
  
  for (let i = 0; i < nums.length; i++) { // Loop dari awal sampai akhir
    const num = nums[i];
    const complement = target - num;    // Hitung angka yang dibutuhkan (complement)

    if (numMap.has(complement)) {     // Cek di map, ada gak complement-nya?
      return [numMap.get(complement), i]; // Kalo ada, BALIKIN INDEKSNYA. (numMap.get(complement) adalah indeks si complement)
    }
    
    numMap.set(num, i);  // Kalo gak ada, simpen num dan indeksnya ke map
  }
  
  return []; // Kalo gak ketemu, balikin array kosong (sesuai soal)
};

console.log(twoSum([2, 7, 11, 15], 9));


