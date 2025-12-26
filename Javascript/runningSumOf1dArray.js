var runningSum = function(nums) {
    const result = new Array(nums.length);

  
    result[0] = nums[0];  // elemen pertama

    for (let i = 1; i < nums.length; i++) {
        result[i] = result[i - 1] + nums[i]; // akumulasi dari nilai sebelumnya
    }

    return result;
};

console.log(runningSum([1, 2, 3, 4]));
