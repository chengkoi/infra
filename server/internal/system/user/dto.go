package user

// CreateUserReq 创建用户请求
type CreateUserReq struct {
	Username string `json:"username" binding:"required" validate:"required,min=3,max=64"`
	Password string `json:"password" binding:"required" validate:"required,min=6,max=32"`
	Nickname string `json:"nickname" binding:"required" validate:"required,max=64"`
	Email    string `json:"email" validate:"omitempty,email,max=128"`
	Phone    string `json:"phone" validate:"omitempty,max=20"`
	Gender   int    `json:"gender" validate:"omitempty,oneof=0 1 2"`
	Remark   string `json:"remark" validate:"omitempty,max=256"`
}

// UpdateUserReq 更新用户请求
type UpdateUserReq struct {
	Nickname string `json:"nickname" validate:"omitempty,max=64"`
	Email    string `json:"email" validate:"omitempty,email,max=128"`
	Phone    string `json:"phone" validate:"omitempty,max=20"`
	Gender   int    `json:"gender" validate:"omitempty,oneof=0 1 2"`
	Status   *int   `json:"status" validate:"omitempty,oneof=0 1"`
	Remark   string `json:"remark" validate:"omitempty,max=256"`
}

// UserQueryReq 用户查询请求
type UserQueryReq struct {
	Page     int    `json:"page" form:"page" validate:"omitempty,min=1"`
	PageSize int    `json:"pageSize" form:"pageSize" validate:"omitempty,min=1,max=100"`
	Username string `json:"username" form:"username"`
	Nickname string `json:"nickname" form:"nickname"`
	Status   *int   `json:"status" form:"status" validate:"omitempty,oneof=0 1"`
}

// UserResp 用户响应
type UserResp struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Gender    int    `json:"gender"`
	Status    int    `json:"status"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// UserListResp 用户列表响应
type UserListResp struct {
	List     []UserResp `json:"list"`
	Total    int64      `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
}

// LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}