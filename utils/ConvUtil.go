package utils

import (
	"github.com/djimenez/iconv-go"
	"strconv"
)

func ConvGB2312ToUTF8(text string) string {
	s, _ := iconv.ConvertString(text, "gb2312", "utf8")
	return s
}

func ConvStringToInt(text string) int {
	t, _ := strconv.Atoi(text)
	return t
}