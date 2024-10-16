package str

import (
	"regexp"
	"strings"
)

// Capitalize 将字符串的首字母转换为大写
func Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	// 将首字母转换为大写，其余部分保持不变
	return strings.ToUpper(string(str[0])) + str[1:]
}

// Lower 将字符串的首字母转换为小写
func Lower(str string) string {
	if len(str) == 0 {
		return str
	}
	// 将首字母转换为大写，其余部分保持不变
	return strings.ToLower(string(str[0])) + str[1:]
}

// ToSnakeCase 将驼峰字符串转换为下划线字符串
func ToSnakeCase(str string) string {
	// 使用正则表达式将驼峰字符串分割成单词
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(str, "${1}_${2}")
	// 将字符串转换为小写
	return strings.ToLower(snake)
}
