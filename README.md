# Go 包导入功能测试项目

这个项目演示了 Go 语言中的包导入功能，包括主目录包、子目录包的导入，以及多模块项目结构。

## 项目结构

```
gomulti/
├── go.mod                      # 主模块定义文件
├── main.go                     # 主程序文件
├── README.md                   # 项目说明文件
├── utils/                      # 工具包
│   ├── utils.go               # 工具函数
│   └── utils_test.go          # 工具函数测试
├── math/                       # 数学运算包
│   └── math.go                # 数学运算函数
├── models/                     # 数据模型包
│   └── user.go                # 用户模型和配置
├── services/                   # 服务包
│   └── database/               # 数据库服务包
│       └── db.go              # 数据库模拟实现
└── tools/                      # 独立子模块
    ├── go.mod                 # 子模块定义文件
    ├── main.go                # 子模块主程序
    ├── README.md              # 子模块文档
    └── fileutils/             # 文件工具包
        └── fileutils.go       # 文件操作函数
```

## 包导入方式演示

### 1. 同级目录包导入

```go
import "github.com/Andrewmatilde/gomulti/utils"
import "github.com/Andrewmatilde/gomulti/math"
import "github.com/Andrewmatilde/gomulti/models"
```

### 2. 嵌套目录包导入

```go
import "github.com/Andrewmatilde/gomulti/services/database"
```

### 3. 子模块包导入

```go
import "github.com/Andrewmatilde/gomulti/tools/fileutils"
```

### 4. 包间依赖

`services/database` 包导入并使用 `models` 包中定义的结构体。

## 功能特性

### utils 包
- `ReverseString(s string) string` - 反转字符串
- `Greet(name string) string` - 生成问候语
- `IsEmpty(s string) bool` - 检查字符串是否为空
- `Contains(s, substr string) bool` - 检查字符串是否包含子字符串

### math 包
- `Add(a, b int) int` - 加法运算
- `Subtract(a, b int) int` - 减法运算
- `Multiply(a, b int) int` - 乘法运算
- `Divide(a, b int) (int, error)` - 除法运算（带错误处理）
- `Max(a, b int) int` - 返回最大值
- `Min(a, b int) int` - 返回最小值
- `Abs(n int) int` - 返回绝对值

### models 包
- `User` 结构体 - 用户数据模型
- `NewUser(name string, age int) *User` - 创建新用户
- `Config` 结构体 - 配置数据模型
- `GetDefaultConfig() *Config` - 获取默认配置

### services/database 包
- `Database` 接口 - 数据库操作接口
- `MockDB` 结构体 - 模拟数据库实现
- 提供用户数据的 CRUD 操作

## 子模块: tools

`tools` 目录是一个独立的子模块，拥有自己的 `go.mod` 文件。

### tools 模块特性
- **独立模块**: 有自己的模块路径 `github.com/Andrewmatilde/gomulti/tools`
- **命令行工具**: 提供文件操作和文本处理功能
- **库包**: 包含 `fileutils` 包，可被其他项目导入使用

### 使用 tools 子模块

```bash
# 进入子模块目录
cd tools

# 使用命令行工具
go run main.go filecount ../
go run main.go fileext ../
go run main.go reverse "Hello World"
go run main.go info

# 构建可执行文件
go build -o tools main.go
```

## 运行项目

### 运行主程序
```bash
go run main.go
```

### 运行测试
```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./utils

# 运行测试并显示详细信息
go test -v ./utils
```

### 构建可执行文件
```bash
go build -o gomulti main.go
./gomulti
```

## 多模块项目的特点

1. **主模块**: `github.com/Andrewmatilde/gomulti`
   - 包含主要的业务逻辑包
   - 演示基本的包导入功能

2. **子模块**: `github.com/Andrewmatilde/gomulti/tools`
   - 独立的模块，有自己的 go.mod
   - 可以独立开发、测试和分发
   - 提供专门的工具功能

## 包导入要点

1. **模块路径**: 每个模块都有自己的模块路径
2. **包名**: 每个包的包名通常与目录名相同
3. **导出符号**: 只有首字母大写的函数、变量、结构体等才能被其他包导入使用
4. **包导入路径**: 使用完整的模块路径 + 包路径进行导入
5. **包间依赖**: 包可以相互依赖，但要避免循环依赖
6. **子模块**: 子模块可以独立管理依赖和版本

## 测试说明

项目包含了完整的测试套件，演示了如何为不同的包编写单元测试。测试文件命名规则：
- 测试文件名以 `_test.go` 结尾
- 测试函数名以 `Test` 开头
- 使用 `testing` 包进行测试

这个项目是学习 Go 包导入机制和多模块项目结构的完整示例，涵盖了从基础到高级的各种使用场景。
