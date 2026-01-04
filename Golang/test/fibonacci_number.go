package test

// func Fib(n int) int { // 0 1 1 2 3 5 8 13 --> 3
// 	if n <= 0 {
// 		return 0
// 	}

// 	if n == 1 {
// 		return 1
// 	}

// 	return Fib(n - 1) + Fib(n - 2)
// }

/* n = 2
fib(1) + fib(0)
1 + 0
*/

// dengan memoization supaya dia O(n) bukan O(2^n)

// func Fib(n int) int {
// 	memo := make(map[int]int)

// 	var dfs func(int)int

// 	dfs = func(n int) int {
// 		if n <= 0 { // base case 1
// 			return 0
// 		}

// 		if n == 1 { // base case 2
// 			return 1
// 		}

// 		if val, found := memo[n]; found { // cek apakah bilangan n udah pernah dihitung
// 			return val
// 		}

// 		memo[n] = dfs(n-1) + dfs(n-2) // kalo belom lakuin recurssion-nya

// 		return memo[n]
// 	}

// 	return dfs(n)
// }

// fibonacci dengan iteration -> o(1) artinya konstan alias performanya paling efisien
func Fib(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	prev := 0
	curr := 1

	for i := 2; i <= n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}

	return curr
}
