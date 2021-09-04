package utils

import "strings"

// SplitThenTrim 分割字符串，然后去除每个子串首尾的空格
func SplitThenTrim(str string, sep string) []string {
	parts := strings.Split(str, sep)
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}
