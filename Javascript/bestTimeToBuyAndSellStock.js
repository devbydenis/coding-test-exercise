var maxProfit = function(prices) {
    let minPrice = prices[0]; // Harga beli termurah (awal)
    let maxProfit = 0;        // Profit maksimum awal = 0

    for (let i = 1; i < prices.length; i++) { // Loop dari hari ke-2
        if (prices[i] < minPrice) {
            minPrice = prices[i]; // Jika ketemu harga lebih murah, update harga beli
        } else {
            const profit = prices[i] - minPrice; // Hitung profit jika dijual hari ini

            if (profit > maxProfit) { // Update profit maksimum
                maxProfit = profit;
            }
        }
    }

    return maxProfit; // Return profit terbesar
};

console.log(maxProfit([7,1,5,3,6,4]));