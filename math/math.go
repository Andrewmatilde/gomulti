package math

import "fmt"

// Add 加法运算
func Add(a, b int) int {
	return a + b
}

// Subtract 减法运算
func Subtract(a, b int) int {
	return a - b
}

// Multiply 乘法运算
func Multiply(a, b int) int {
	return a * b
}

// Divide 除法运算
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// Max 返回两个数中的最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min 返回两个数中的最小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Abs 返回绝对值
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
