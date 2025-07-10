package models

import "fmt"

// User 用户结构体
type User struct {
	Name string
	Age  int
}

// NewUser 创建新用户
func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

// String 返回用户的字符串表示
func (u *User) String() string {
	return fmt.Sprintf("用户: %s, 年龄: %d", u.Name, u.Age)
}

// IsAdult 检查用户是否成年
func (u *User) IsAdult() bool {
	return u.Age >= 18
}

// Config 配置结构体
type Config struct {
	Host  string
	Port  int
	Debug bool
}

// GetDefaultConfig 获取默认配置
func GetDefaultConfig() *Config {
	return &Config{
		Host:  "localhost",
		Port:  8080,
		Debug: true,
	}
}

// String 返回配置的字符串表示
func (c *Config) String() string {
	return fmt.Sprintf("Host: %s, Port: %d, Debug: %t", c.Host, c.Port, c.Debug)
}
