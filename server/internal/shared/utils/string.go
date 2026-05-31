package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
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

// RandomString 生成随机字符串
func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}
	return string(b)
}