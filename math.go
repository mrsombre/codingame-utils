package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
