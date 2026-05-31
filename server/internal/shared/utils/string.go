package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// MD5 MD5加密
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// IsEmpty 判断字符串是否为空
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Contains 判断字符串是否包含
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// ToUpper 转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}