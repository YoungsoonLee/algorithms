package main

import "fmt"

func UniquePath(m, n int) int {
	// init
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	fmt.Println(dp)

	// time: O(m*n)
	// space: O(m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	m, n := 2, 3
	fmt.Println(UniquePath(m, n))
}
