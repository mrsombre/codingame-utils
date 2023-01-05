package main

import (
	"strconv"
)

func intToStr(x int) string {
	return strconv.Itoa(x)
}

func boolToInt(x bool) int {
	if x {
		return 1
	}
	return 0
}
