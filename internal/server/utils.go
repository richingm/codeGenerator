package server

import (
	"fmt"
	"os"
	"path/filepath"
)

// createAndWriteFile 创建并且写入文件
func createAndWriteFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating or truncating the file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing content to the file: %v", err)
	}

	fmt.Printf("File written: %s\n", filePath)
	return nil
}

func CreateFile(basePath, nestedPath, fileName, content string) error {
	fullDirPath := filepath.Join(basePath, nestedPath)
	if _, err := os.Stat(fullDirPath); os.IsNotExist(err) {
		err := os.MkdirAll(fullDirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	fullFilePath := filepath.Join(fullDirPath, fileName)
	err := createAndWriteFile(fullFilePath, content)
	if err != nil {
		return err
	}
	return nil
}
