package operlog

import "time"

// OperLog 操作日志实体
type OperLog struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:日志ID"`
	UserID    uint      `json:"user_id" gorm:"index;comment:用户ID"`
	Username  string    `json:"username" gorm:"size:64;index;comment:用户名"`
	Module    string    `json:"module" gorm:"size:64;comment:操作模块"`
	Action    string    `json:"action" gorm:"size:128;comment:操作类型"`
	Method    string    `json:"method" gorm:"size:10;comment:请求方法"`
	Path      string    `json:"path" gorm:"size:256;comment:请求路径"`
	IP        string    `json:"ip" gorm:"size:64;comment:IP地址"`
	UserAgent string    `json:"user_agent" gorm:"size:512;comment:UserAgent"`
	Request   string    `json:"request" gorm:"type:text;comment:请求参数"`
	Response  string    `json:"response" gorm:"type:text;comment:响应内容"`
	Duration  int64     `json:"duration" gorm:"comment:耗时(ms)"`
	Status    int       `json:"status" gorm:"default:1;comment:状态 1成功 0失败"`
	CreatedAt time.Time `json:"created_at"`
}

func (OperLog) TableName() string {
	return "sys_oper_log"
}
