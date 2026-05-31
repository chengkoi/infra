package user

import "time"

// User 用户实体
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;comment:用户ID"`
	Username  string    `json:"username" gorm:"uniqueIndex;size:64;not null;comment:用户名"`
	Password  string    `json:"-" gorm:"size:128;not null;comment:密码"`
	Nickname  string    `json:"nickname" gorm:"size:64;comment:昵称"`
	Email     string    `json:"email" gorm:"size:128;comment:邮箱"`
	Phone     string    `json:"phone" gorm:"size:20;comment:手机号"`
	Avatar    string    `json:"avatar" gorm:"size:256;comment:头像"`
	Gender    int       `json:"gender" gorm:"default:0;comment:性别 0未知 1男 2女"`
	Status    int       `json:"status" gorm:"default:1;comment:状态 0禁用 1启用"`
	Remark    string    `json:"remark" gorm:"size:256;comment:备注"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 表名
func (User) TableName() string {
	return "sys_user"
}