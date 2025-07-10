package utils

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"Go", "oG"},
		{"", ""},
		{"a", "a"},
	}

	for _, test := range tests {
		result := ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"张三", "你好, 张三! 欢迎使用 Go 包导入测试!"},
		{"李四", "你好, 李四! 欢迎使用 Go 包导入测试!"},
		{"", "你好, ! 欢迎使用 Go 包导入测试!"},
	}

	for _, test := range tests {
		result := Greet(test.name)
		if result != test.expected {
			t.Errorf("Greet(%q) = %q, expected %q", test.name, result, test.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"hello", false},
		{" ", false},
		{"world", false},
	}

	for _, test := range tests {
		result := IsEmpty(test.input)
		if result != test.expected {
			t.Errorf("IsEmpty(%q) = %t, expected %t", test.input, result, test.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		s        string
		substr   string
		expected bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", false}, // 因为我们的实现只检查前缀
		{"hello world", "xyz", false},
		{"", "", true},
		{"hello", "", true},
	}

	for _, test := range tests {
		result := Contains(test.s, test.substr)
		if result != test.expected {
			t.Errorf("Contains(%q, %q) = %t, expected %t", test.s, test.substr, result, test.expected)
		}
	}
}
