package tools

import "strconv"

func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
