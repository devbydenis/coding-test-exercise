function PalindromNumber(number) {
  const reversedNumber = number.toString().split("").reverse().join("")

  return reversedNumber == number ? true : false
}

console.time()
console.log(PalindromNumber(121))
console.log(PalindromNumber(1123197698))
console.timeEnd()