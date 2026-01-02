function containDuplicate(nums) {
  const numbersMap = new Set();

  for (const num of nums) {
    if (numbersMap.has(num)) {
      return true;
    }

    numbersMap.add(num);
  }

  return false;
}

console.time()
console.log(containDuplicate([1, 2, 3, 4]));
console.log(containDuplicate([1, 2, 3, 4, 1]));
console.timeEnd()
