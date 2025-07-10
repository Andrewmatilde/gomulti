package main

import (
	"fmt"

	"github.com/Andrewmatilde/gomulti/math"
	"github.com/Andrewmatilde/gomulti/models"
	"github.com/Andrewmatilde/gomulti/services/database"
	"github.com/Andrewmatilde/gomulti/utils"
)

func main() {
	fmt.Println("=== Go 包导入功能测试 ===")

	// 测试 utils 包
	fmt.Println("\n1. 测试 utils 包:")
	result := utils.ReverseString("Hello World")
	fmt.Printf("字符串反转: %s\n", result)

	greeting := utils.Greet("张三")
	fmt.Printf("问候语: %s\n", greeting)

	isEmpty := utils.IsEmpty("")
	fmt.Printf("空字符串检查: %t\n", isEmpty)

	// 测试 math 包
	fmt.Println("\n2. 测试 math 包:")
	sum := math.Add(10, 20)
	fmt.Printf("加法: 10 + 20 = %d\n", sum)

	product := math.Multiply(5, 6)
	fmt.Printf("乘法: 5 * 6 = %d\n", product)

	maxVal := math.Max(15, 25)
	fmt.Printf("最大值: Max(15, 25) = %d\n", maxVal)

	// 测试 models 包
	fmt.Println("\n3. 测试 models 包:")
	user1 := models.NewUser("李四", 25)
	fmt.Printf("用户信息: %s\n", user1.String())
	fmt.Printf("是否成年: %t\n", user1.IsAdult())

	user2 := models.NewUser("王五", 16)
	fmt.Printf("用户信息: %s\n", user2.String())
	fmt.Printf("是否成年: %t\n", user2.IsAdult())

	// 测试配置
	config := models.GetDefaultConfig()
	fmt.Printf("默认配置: %s\n", config.String())

	// 测试嵌套包导入 (services/database)
	fmt.Println("\n4. 测试嵌套包导入 (services/database):")
	db := database.NewMockDB()

	// 连接数据库
	if err := db.Connect(); err != nil {
		fmt.Printf("连接错误: %v\n", err)
	}

	// 保存用户
	if err := db.SaveUser(user1); err != nil {
		fmt.Printf("保存用户错误: %v\n", err)
	}

	if err := db.SaveUser(user2); err != nil {
		fmt.Printf("保存用户错误: %v\n", err)
	}

	// 获取用户数量
	userCount := db.GetUserCount()
	fmt.Printf("数据库中用户数量: %d\n", userCount)

	// 获取所有用户
	allUsers := db.GetAllUsers()
	fmt.Println("数据库中所有用户:")
	for i, u := range allUsers {
		fmt.Printf("  %d. %s\n", i+1, u.String())
	}

	// 关闭数据库
	if err := db.Close(); err != nil {
		fmt.Printf("关闭数据库错误: %v\n", err)
	}

	// 测试包间依赖
	fmt.Println("\n5. 测试包间依赖:")
	fmt.Printf("用户名长度: %d\n", len(user1.Name))
	reversedName := utils.ReverseString(user1.Name)
	fmt.Printf("反转用户名: %s\n", reversedName)

	ageSquared := math.Multiply(user1.Age, user1.Age)
	fmt.Printf("年龄的平方: %d\n", ageSquared)

	fmt.Println("\n=== 包导入测试完成 ===")
}
