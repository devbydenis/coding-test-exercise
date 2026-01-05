function lengthOfLongestSubstring(str) { // p w w k e w
  let charMap = new Map() // p: 0
  let left = 0
  let maxLength = 0

  for (let right = 0; right < str.length; right++) {
    const char = str[right]

    // cek apakah character yang sedang di looping itu duplicate -> shrink
    if (charMap.has(char)) {
      const lastIndex = charMap.get(char)

      if (lastIndex + 1 > left) {
        left = lastIndex + 1
      }
    }

    // set ke Map karna unique -> expand
    charMap.set(char, right)

    // panjang karakter yang ada didalam window saat ini
    let currentLength = right - left + 1

    if (currentLength > maxLength) {
      maxLength = currentLength
    }
  }

  return maxLength
}

console.time()
console.log(lengthOfLongestSubstring("bbbbbbb"))
console.log(lengthOfLongestSubstring("pwwkew"))
console.timeEnd()
