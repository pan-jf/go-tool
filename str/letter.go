package str

import "strings"

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
