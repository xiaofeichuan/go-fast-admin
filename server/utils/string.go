package utils

import "strings"

// ToPascalCase
// 蛇形命名法=>帕斯卡命名法
// 例如：user_name->UserName
func ToPascalCase(snakeCase string) string {
	var result string
	nameList := strings.Split(snakeCase, "_")
	for i := 0; i < len(nameList); i++ {
		name := nameList[i]
		// 第一个字母
		firstChar := string([]byte(name)[:1])
		// 剩余字母
		otherChar := string([]byte(name)[1:])
		result += strings.ToUpper(firstChar) + otherChar
	}
	return result
}

// ToCamelCase
// 蛇形命名法=>驼峰式命名法
// 例如：user_name->userName
func ToCamelCase(snakeCase string) string {
	var result string
	nameList := strings.Split(snakeCase, "_")
	for i := 0; i < len(nameList); i++ {
		name := nameList[i]
		if i == 0 {
			// 第一个单词
			result += strings.ToLower(name)
		} else {
			// 第一个字母
			firstChar := string([]byte(name)[:1])
			// 剩余字母
			otherChar := string([]byte(name)[1:])
			result += strings.ToUpper(firstChar) + otherChar
		}
	}
	return result
}

// GetGoType 获取go的数据类型
// mysql数据类型=>go数据类型
func GetGoType(columnName string, columnType string) string {
	if strings.Contains(columnType, "int") {
		if columnName == "id" || strings.Contains(columnType, "_id") {
			return "uint64"
		}
		//判断是不是bool
		var isStr = string([]byte(columnName)[:3])
		if columnName == isStr {
			return "bool"
		}
		return "int"
	} else if strings.Contains(columnType, "datetime") {
		return "time.Time"
	} else if strings.Contains(columnType, "char") {
		return "string"
	} else if strings.Contains(columnType, "decimal") {
		//需要扩展
		return "decimal.Decimal"
	} else {
		return columnType
	}
}
