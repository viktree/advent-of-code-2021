package lib

import "strconv"

func ToInt(in string) int {
	if num, err := strconv.Atoi(in); err != nil {
		panic(err)
	} else {
		return num
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
