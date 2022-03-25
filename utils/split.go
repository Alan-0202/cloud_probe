package utils

import "strings"

func SplitPy(s string) string {  // 1.py get 1
	return strings.Split(s, ".")[0]
}
