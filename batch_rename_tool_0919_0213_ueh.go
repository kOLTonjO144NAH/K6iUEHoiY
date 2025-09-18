// 代码生成时间: 2025-09-19 02:13:02
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

// RenamePattern 是批量重命名的模式结构体
type RenamePattern struct {
    Search string // 要搜索的字符串
    Replace string // 替换搜索字符串的内容
}

// RenameBatch 执行批量文件重命名
func RenameBatch(directory string, renamePatterns []RenamePattern) error {
    files, err := os.ReadDir(directory)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            continue // 跳过子目录
        }

        fileName := file.Name()
        for _, pattern := range renamePatterns {
            if strings.Contains(fileName, pattern.Search) {
                newFileName := strings.ReplaceAll(fileName, pattern.Search, pattern.Replace)
                if err := os.Rename(filepath.Join(directory, fileName), filepath.Join(directory, newFileName)); err != nil {
                    return fmt.Errorf("failed to rename file %s to %s: %w", fileName, newFileName, err)
                }
                fmt.Printf("Renamed '%s' to '%s'
", fileName, newFileName)
                break // 找到一个匹配的模式后停止搜索
            }
        }
    }

    return nil
}

func main() {
    // 示例：将目录下所有文件中的'old'字符串替换为'new'
    directory := "./example_directory"
    renamePatterns := []RenamePattern{{"old", "new"}}

    if err := RenameBatch(directory, renamePatterns); err != nil {
        fmt.Printf("Error: %s
", err)
    } else {
        fmt.Println("Batch renaming completed successfully.")
    }
}
