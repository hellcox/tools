package util

import (
	"regexp"
	"strings"
)

// 提取字符串中的数字
func GetNum(str string) string {
	patten := regexp.MustCompile(`\d`)
	newStr := strings.Join(patten.FindAllString(str, -1), "")
	return newStr
}
