function validPalindrome(words) {
  let left = 0;
  let right = words.length - 1;

  while (left < right) {
    while(left < right && !isAlpanumeric(words[left])) {
      left ++
    }

    while(left < right && !isAlpanumeric(words[right])) {
      right --
    }

    if (words[left].toLowerCase() !== words[right].toLowerCase()) {
      return false
    }

    left++
    right--
  }

  return true
}

console.time()
console.log(validPalindrome("kasur rusak"))
console.log(validPalindrome("bahlil bangsat"))
console.timeEnd()

function isAlpanumeric(char) {
  return /^[a-z0-9]$/i.test(char)
}
