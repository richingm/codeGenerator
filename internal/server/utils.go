package server

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
	"unicode"
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

	return nil
}

func CreateFile(fullDirPath, fileName, content string) error {
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

func CombineWords(name string, staff string) string {
	return fmt.Sprintf("%s%s", name, staff)
}

func lowercaseFirstLetter(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

// underscoreToCamelCase 下划线转驼峰,首字母大写
func underscoreToCamelCase(s string) string {
	// Split the string into words on underscores
	words := strings.Split(s, "_")
	// Capitalize the first letter of each word and join them back together
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

// lowerCamelCase 首字母转小写
func lowerCamelCase(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func generateStructFromTable(repoPo string, columns []gorm.ColumnType) string {
	res := "type " + repoPo + " struct {\n"
	for _, column := range columns {
		goType := getGoType(column.DatabaseTypeName())
		fieldName := formatFieldName(column.Name())
		res += fmt.Sprintf("\t%s %s `gorm:\"column:%s\"`\n", fieldName, goType, column.Name())
	}
	res += "}"
	return res
}

func getGoType(dbType string) string {
	// 这里是一个简单的映射，可能需要根据实际情况调整
	switch {
	case strings.Contains(dbType, "int8"):
		return "int64"
	case strings.Contains(dbType, "bigint"):
		return "int64"
	case strings.Contains(dbType, "integer"):
		return "int64"
	case strings.Contains(dbType, "text"),
		strings.Contains(dbType, "varchar"):
		return "string"
	case strings.Contains(dbType, "boolean"):
		return "bool"
	case strings.Contains(dbType, "timestamp"),
		strings.Contains(dbType, "date"):
		return "time.Time"
	default:
		return "string" // 默认情况下，使用string类型
	}
}

func formatFieldName(columnName string) string {
	return strings.Replace(strings.Title(strings.Replace(columnName, "_", " ", -1)), " ", "", -1)
}
