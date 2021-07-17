package utils

import (
	"fmt"
	"strings"
)

func DefaultToString(item interface{}) string {
	if item == nil {
		return "NULL"
	}
	if str, ok := item.(string); ok {
		if str == "" {
			return "NULL"
		}
		return str
	}
	return fmt.Sprintf("%v", item) // type number
}

func IsEmpty(item interface{}) bool {
	if item == nil {
		return true
	}
	if str, ok := item.(string); ok {
		if removeWhiteSpaces(str) == "" {
			return true
		}
	}
	return false
}

func removeWhiteSpaces(str string) string {
	return strings.Replace(str, " ", "", -1)
}
