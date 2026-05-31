package loginlog

import "time"

// LoginLog 登录日志实体
type LoginLog struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:日志ID"`
	UserID    uint      `json:"user_id" gorm:"index;comment:用户ID"`
	Username  string    `json:"username" gorm:"size:64;index;comment:用户名"`
	IP        string    `json:"ip" gorm:"size:64;comment:IP地址"`
	UserAgent string    `json:"user_agent" gorm:"size:512;comment:UserAgent"`
	Status    int       `json:"status" gorm:"default:1;comment:状态 1成功 0失败"`
	Message   string    `json:"message" gorm:"size:256;comment:消息"`
	CreatedAt time.Time `json:"created_at"`
}

func (LoginLog) TableName() string {
	return "sys_login_log"
}
