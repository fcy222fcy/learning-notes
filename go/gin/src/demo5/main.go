package main

import "fmt"

func main() {
	//router := gin.Default()
	//
	//router.GET("/hello")
	fmt.Println(tribonacci(4))
}
func tribonacci(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		switch {
		case i == 0:
			dp[i] = 0
		case i == 1:
			dp[i] = 1
		case i == 2:
			dp[i] = 1
		default:
			dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
			fmt.Printf("dp[%v] = %v\n", i, dp[i])
		}
	}
	return dp[n]
}
