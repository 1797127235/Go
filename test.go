package main


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPower(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	i := 0
	ans := 1

	for i < n {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		ans = max(ans, j-i)
		i = j
	}

	return ans
}
