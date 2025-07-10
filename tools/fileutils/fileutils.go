package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileInfo 文件信息结构体
type FileInfo struct {
	Name      string
	Size      int64
	Extension string
	IsDir     bool
}

// CountFiles 统计目录中的文件数量
func CountFiles(dir string) (int, error) {
	count := 0
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			count++
		}
		return nil
	})
	return count, err
}

// GetFileExtensions 获取目录中所有文件的扩展名
func GetFileExtensions(dir string) ([]string, error) {
	extensions := make(map[string]bool)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if ext != "" {
				extensions[ext] = true
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(extensions))
	for ext := range extensions {
		result = append(result, ext)
	}

	return result, nil
}

// GetFileList 获取目录中所有文件的详细信息
func GetFileList(dir string) ([]FileInfo, error) {
	var files []FileInfo

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileInfo := FileInfo{
			Name:      info.Name(),
			Size:      info.Size(),
			Extension: strings.ToLower(filepath.Ext(path)),
			IsDir:     info.IsDir(),
		}

		files = append(files, fileInfo)
		return nil
	})

	return files, err
}

// CreateDirectoryIfNotExists 如果目录不存在则创建
func CreateDirectoryIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("创建目录失败: %v", err)
		}
		fmt.Printf("目录 %s 创建成功\n", dir)
	}
	return nil
}

// GetFileSize 获取文件大小（字节）
func GetFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

// FormatFileSize 格式化文件大小显示
func FormatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
