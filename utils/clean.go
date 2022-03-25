package utils

import "strings"

func CleanStr(s string) string{
	str := strings.Replace(s, "\r", "", -1 )
	str = strings.Replace(str, "\n", "", -1)
	return str
}