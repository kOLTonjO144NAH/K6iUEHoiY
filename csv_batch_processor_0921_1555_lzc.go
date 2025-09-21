// 代码生成时间: 2025-09-21 15:55:48
package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// ProcessCSVFile is a function that processes a single CSV file.
func ProcessCSVFile(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()

    reader := csv.NewReader(bufio.NewReader(file))
    records, err := reader.ReadAll()
    if err != nil {
        return fmt.Errorf("failed to read CSV file: %w", err)
    }

    // Process each record (for demonstration, we just print it)
    for _, record := range records {
        fmt.Println(record)
    }
    return nil
}

// ProcessCSVFilesInDirectory is a function that processes all CSV files in a given directory.
func ProcessCSVFilesInDirectory(directoryPath string) error {
    // Get a list of files in the directory
    files, err := os.ReadDir(directoryPath)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        if !file.IsDir() && strings.HasSuffix(file.Name(), ".csv") {
            filePath := filepath.Join(directoryPath, file.Name())
            if err := ProcessCSVFile(filePath); err != nil {
                log.Printf("Error processing file %s: %v", filePath, err)
                continue // Continue processing other files
            }
        }
    }
    return nil
}

// Main function to start the CSV batch processing.
func main() {
    directoryPath := "./csv_files" // The directory containing CSV files
    if err := ProcessCSVFilesInDirectory(directoryPath); err != nil {
        log.Fatalf("Failed to process CSV files: %v", err)
    } else {
        fmt.Println("CSV files processed successfully.")
    }
}
