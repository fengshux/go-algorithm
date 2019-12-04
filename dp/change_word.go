package dp

func ChangeWord(word1, word2 string)  int{
	n1 := len(word1)
	n2 := len(word2)
	dp := make([][]int, n1 + 1)
	for i := range dp {
		dp[i] = make([]int, n2 +1)
	}
	for j:=1; j <= n2; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	for i:=1; i <= n1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for i := 1; i <= n1; i++ {
		for j:=1; j<= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = minThree(dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]) + 1
			}
		}
	}	
	return dp[n1][n2]
}


func minThree(a, b, c int)  (min int ){	
	if a < b {
		min = a
	} else {
		min = b
	}
	if min > c {
		min = c
	}
	return 
}
