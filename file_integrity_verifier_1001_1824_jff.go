// 代码生成时间: 2025-10-01 18:24:35
package main

import (
    "crypto/sha256"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
)

// FileIntegrityVerifier 结构体用于文件完整性校验
type FileIntegrityVerifier struct {
    // 存储文件的期望散列值
    ExpectedHashes map[string]string
}

// NewFileIntegrityVerifier 创建一个新的文件完整性校验器
func NewFileIntegrityVerifier() *FileIntegrityVerifier {
    return &FileIntegrityVerifier{
        ExpectedHashes: make(map[string]string),
    }
}

// AddExpectedHash 添加一个文件的期望散列值
func (verifier *FileIntegrityVerifier) AddExpectedHash(filePath, hash string) {
    verifier.ExpectedHashes[filePath] = hash
}

// VerifyAll 校验所有添加的文件的完整性
func (verifier *FileIntegrityVerifier) VerifyAll() ([]string, error) {
    failedFiles := []string{}
    for filePath, expectedHash := range verifier.ExpectedHashes {
        actualHash, err := verifier.hashFile(filePath)
        if err != nil {
            return nil, err
        }
        if actualHash != expectedHash {
            failedFiles = append(failedFiles, filePath)
        }
    }
    return failedFiles, nil
}

// hashFile 计算给定文件的SHA256散列值
func (verifier *FileIntegrityVerifier) hashFile(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hash := sha256.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", err
    }
    return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func main() {
    verifier := NewFileIntegrityVerifier()
    verifier.AddExpectedHash("./example.txt", "expected_hash_value")
    failedFiles, err := verifier.VerifyAll()
    if err != nil {
        fmt.Println("Error verifying files: ", err)
    } else if len(failedFiles) > 0 {
        fmt.Println("Files with failed integrity checks: ", failedFiles)
    } else {
        fmt.Println("All files verified successfully.")
    }
}
