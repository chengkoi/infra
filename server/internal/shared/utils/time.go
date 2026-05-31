package utils

import "time"

// FormatTime 格式化时间
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FormatDate 格式化日期
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// NowTimestamp 获取当前时间戳(秒)
func NowTimestamp() int64 {
	return time.Now().Unix()
}

// NowMilliTimestamp 获取当前时间戳(毫秒)
func NowMilliTimestamp() int64 {
	return time.Now().UnixMilli()
}