package common

import "strconv"

func ResponseBodyWithMessage(msg string, dataName string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		dataName:  data,
	}
}

func ResponseBodyWithError(errorMsg string) map[string]interface{} {
	return map[string]interface{}{
		"error": errorMsg,
	}
}

func StringToUint(s string) (uint, error) {
	num, err := strconv.ParseUint(s, 10, 64)
	return uint(num), err
}

func StringToInt(s string) (int, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	return int(num), err
}

func UintToString(i uint) string {
	i64 := uint64(i)
	return strconv.FormatUint(i64, 10)
}
