package loginlog

// LoginLogQueryReq 登录日志查询请求
type LoginLogQueryReq struct {
	Page     int    `json:"page" form:"page" validate:"omitempty,min=1"`
	PageSize int    `json:"pageSize" form:"pageSize" validate:"omitempty,min=1,max=100"`
	Username string `json:"username" form:"username"`
	Status   *int   `json:"status" form:"status" validate:"omitempty,oneof=0 1"`
}

// LoginLogResp 登录日志响应
type LoginLogResp struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

// LoginLogListResp 登录日志列表响应
type LoginLogListResp struct {
	List     []LoginLogResp `json:"list"`
	Total    int64          `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}
