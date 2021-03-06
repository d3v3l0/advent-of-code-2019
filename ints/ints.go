package ints

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func Max64(x, y int64) int64 {
	if x >= y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func Gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return Gcd(n, m%n)
}

func Gcd64(m, n int64) int64 {
	if n == 0 {
		return m
	}
	return Gcd64(n, m%n)
}

func Copied64(x []int64) []int64 {
	y := make([]int64, len(x))
	copy(y, x)
	return y
}
