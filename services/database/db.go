package database

import (
	"fmt"

	"github.com/Andrewmatilde/gomulti/models"
)

// Database 数据库接口
type Database interface {
	Connect() error
	Close() error
	SaveUser(user *models.User) error
}

// MockDB 模拟数据库实现
type MockDB struct {
	connected bool
	users     []*models.User
}

// NewMockDB 创建模拟数据库
func NewMockDB() *MockDB {
	return &MockDB{
		connected: false,
		users:     make([]*models.User, 0),
	}
}

// Connect 连接数据库
func (db *MockDB) Connect() error {
	if db.connected {
		return fmt.Errorf("数据库已连接")
	}
	db.connected = true
	fmt.Println("数据库连接成功")
	return nil
}

// Close 关闭数据库连接
func (db *MockDB) Close() error {
	if !db.connected {
		return fmt.Errorf("数据库未连接")
	}
	db.connected = false
	fmt.Println("数据库连接已关闭")
	return nil
}

// SaveUser 保存用户到数据库
func (db *MockDB) SaveUser(user *models.User) error {
	if !db.connected {
		return fmt.Errorf("数据库未连接")
	}
	db.users = append(db.users, user)
	fmt.Printf("用户 %s 已保存到数据库\n", user.Name)
	return nil
}

// GetAllUsers 获取所有用户
func (db *MockDB) GetAllUsers() []*models.User {
	return db.users
}

// GetUserCount 获取用户数量
func (db *MockDB) GetUserCount() int {
	return len(db.users)
}
