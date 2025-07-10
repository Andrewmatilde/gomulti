package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Andrewmatilde/gomulti/tools/fileutils"
)

func main() {
	fmt.Println("=== Go 子模块工具 ===")

	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go <命令>")
		fmt.Println("可用命令:")
		fmt.Println("  filecount <目录>    - 统计目录中的文件数量")
		fmt.Println("  fileext <目录>      - 显示目录中的文件扩展名")
		fmt.Println("  filelist <目录>     - 显示目录中的文件列表")
		fmt.Println("  reverse <文本>      - 反转文本")
		fmt.Println("  info               - 显示模块信息")
		return
	}

	command := os.Args[1]

	switch command {
	case "filecount":
		if len(os.Args) < 3 {
			fmt.Println("请提供目录路径")
			return
		}
		dir := os.Args[2]
		count, err := fileutils.CountFiles(dir)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}
		fmt.Printf("目录 %s 中共有 %d 个文件\n", dir, count)

	case "fileext":
		if len(os.Args) < 3 {
			fmt.Println("请提供目录路径")
			return
		}
		dir := os.Args[2]
		extensions, err := fileutils.GetFileExtensions(dir)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}
		fmt.Printf("目录 %s 中的文件扩展名: %v\n", dir, extensions)

	case "filelist":
		if len(os.Args) < 3 {
			fmt.Println("请提供目录路径")
			return
		}
		dir := os.Args[2]
		files, err := fileutils.GetFileList(dir)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}
		fmt.Printf("目录 %s 中的文件列表:\n", dir)
		for _, file := range files {
			if file.IsDir {
				fmt.Printf("  [DIR]  %s\n", file.Name)
			} else {
				sizeStr := fileutils.FormatFileSize(file.Size)
				fmt.Printf("  [FILE] %s (%s) %s\n", file.Name, sizeStr, file.Extension)
			}
		}

	case "reverse":
		if len(os.Args) < 3 {
			fmt.Println("请提供要反转的文本")
			return
		}
		text := strings.Join(os.Args[2:], " ")
		reversed := reverseString(text)
		fmt.Printf("原文本: %s\n", text)
		fmt.Printf("反转后: %s\n", reversed)

	case "info":
		showModuleInfo()

	default:
		fmt.Printf("未知命令: %s\n", command)
	}
}

// reverseString 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// showModuleInfo 显示模块信息
func showModuleInfo() {
	fmt.Println("模块名称: github.com/Andrewmatilde/gomulti/tools")
	fmt.Println("模块类型: 独立子模块")
	fmt.Println("功能: 提供文件工具和文本处理功能")
	fmt.Println("包含的子包:")
	fmt.Println("  - fileutils: 文件操作工具")
	fmt.Println("这是一个演示 Go 子模块的示例项目")
}
