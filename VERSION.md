# 版本发布说明

## 发布历史

### v1.0.0 (主模块) - 2025-07-10

**主模块**: `github.com/Andrewmatilde/gomulti`

#### 功能特性
- ✅ 完整的包导入功能演示
- ✅ 包含 utils, math, models, services 等核心包
- ✅ 支持同级目录包和嵌套包导入
- ✅ 包含单元测试和详细文档
- ✅ 演示包间依赖关系

#### 使用方法
```bash
# 克隆项目
git clone https://github.com/Andrewmatilde/gomulti.git
cd gomulti

# 切换到 v1.0.0 版本
git checkout v1.0.0

# 运行主程序
go run main.go

# 运行测试
go test ./...
```

#### 导入使用
```go
import (
    "github.com/Andrewmatilde/gomulti/utils"
    "github.com/Andrewmatilde/gomulti/math"
    "github.com/Andrewmatilde/gomulti/models"
    "github.com/Andrewmatilde/gomulti/services/database"
)
```

---

### tools/v1.0.0 (子模块) - 2025-07-10

**子模块**: `github.com/Andrewmatilde/gomulti/tools`

#### 功能特性
- ✅ 独立的 Go 子模块
- ✅ 包含 fileutils 文件操作工具包
- ✅ 提供命令行工具：文件统计、扩展名分析、文本处理
- ✅ 支持作为库导入使用
- ✅ 完整的模块文档和使用说明

#### 使用方法
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

#### 导入使用
```go
import "github.com/Andrewmatilde/gomulti/tools/fileutils"

// 使用文件工具
count, err := fileutils.CountFiles("./")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("文件数量: %d\n", count)
```

## Go 模块使用

### 使用主模块
```bash
# 在你的项目中
go mod init myproject
go get github.com/Andrewmatilde/gomulti@v1.0.0
```

### 使用子模块
```bash
# 在你的项目中
go mod init myproject
go get github.com/Andrewmatilde/gomulti/tools@tools/v1.0.0
```

## 多模块项目结构

这个项目演示了 Go 多模块项目的最佳实践：

1. **主模块标签**: 直接使用版本号 `v1.0.0`
2. **子模块标签**: 使用路径前缀 `tools/v1.0.0`
3. **独立版本控制**: 每个模块可以独立发布版本
4. **灵活依赖**: 可以只依赖需要的特定模块

## 下一步计划

- [ ] 添加更多工具包
- [ ] 完善错误处理
- [ ] 添加性能基准测试
- [ ] 支持更多文件格式
- [ ] 添加配置文件支持

## 反馈和贡献

欢迎通过 GitHub Issues 提交问题和建议！ 