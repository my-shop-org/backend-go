package pkg

import (
	"strconv"
	"unicode"
)

func UintToString(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

func StringToUint(s string) uint {
	u64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint(u64)
}

func CapitalizeFirstLetter(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
