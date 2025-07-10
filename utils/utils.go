package utils

import "fmt"

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Greet 生成问候语
func Greet(name string) string {
	return fmt.Sprintf("你好, %s! 欢迎使用 Go 包导入测试!", name)
}

// IsEmpty 检查字符串是否为空
func IsEmpty(s string) bool {
	return len(s) == 0
}

// Contains 检查字符串是否包含子字符串
func Contains(s, substr string) bool {
	return len(s) >= len(substr) && s[0:len(substr)] == substr
}
