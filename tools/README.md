# Tools 子模块

这是 `github.com/Andrewmatilde/gomulti/tools` 子模块，演示了如何在 Go 项目中创建和使用独立的子模块。

## 模块结构

```
tools/
├── go.mod              # 独立的模块定义
├── main.go             # 主程序（命令行工具）
├── README.md           # 模块文档
└── fileutils/          # 文件工具包
    └── fileutils.go    # 文件操作函数
```

## 功能特性

### 命令行工具
- `filecount <目录>` - 统计目录中的文件数量
- `fileext <目录>` - 显示目录中的文件扩展名
- `filelist <目录>` - 显示目录中的文件列表
- `reverse <文本>` - 反转文本
- `info` - 显示模块信息

### fileutils 包
- `CountFiles(dir string)` - 统计目录中的文件数量
- `GetFileExtensions(dir string)` - 获取目录中所有文件的扩展名
- `GetFileList(dir string)` - 获取目录中所有文件的详细信息
- `CreateDirectoryIfNotExists(dir string)` - 如果目录不存在则创建
- `GetFileSize(filename string)` - 获取文件大小
- `FormatFileSize(size int64)` - 格式化文件大小显示

## 使用方法

### 作为命令行工具使用

```bash
# 统计文件数量
go run main.go filecount ../

# 显示文件扩展名
go run main.go fileext ../

# 显示文件列表
go run main.go filelist ../

# 反转文本
go run main.go reverse "Hello World"

# 显示模块信息
go run main.go info
```

### 作为库使用

```go
import "github.com/Andrewmatilde/gomulti/tools/fileutils"

// 统计文件数量
count, err := fileutils.CountFiles("./")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("文件数量: %d\n", count)

// 获取文件扩展名
extensions, err := fileutils.GetFileExtensions("./")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("扩展名: %v\n", extensions)
```

## 构建可执行文件

```bash
# 构建
go build -o tools main.go

# 使用
./tools filecount ../
./tools info
```

## 子模块的优势

1. **独立管理**: 拥有自己的 go.mod 文件，可以独立版本控制
2. **清晰边界**: 功能模块化，职责分离
3. **可重用**: 可以被其他项目导入使用
4. **灵活部署**: 可以单独构建和分发

## 注意事项

- 子模块有自己的 go.mod 文件
- 子模块可以有自己的依赖管理
- 子模块内部的包可以相互导入
- 外部项目可以导入子模块的包

这个子模块展示了 Go 多模块项目的典型结构和使用方式。 