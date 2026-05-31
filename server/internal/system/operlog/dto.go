package operlog

// OperLogQueryReq 操作日志查询请求
type OperLogQueryReq struct {
	Page     int    `json:"page" form:"page" validate:"omitempty,min=1"`
	PageSize int    `json:"pageSize" form:"pageSize" validate:"omitempty,min=1,max=100"`
	Username string `json:"username" form:"username"`
	Module   string `json:"module" form:"module"`
	Status   *int   `json:"status" form:"status" validate:"omitempty,oneof=0 1"`
}

// OperLogResp 操作日志响应
type OperLogResp struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Module    string `json:"module"`
	Action    string `json:"action"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Duration  int64  `json:"duration"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
}

// OperLogListResp 操作日志列表响应
type OperLogListResp struct {
	List     []OperLogResp `json:"list"`
	Total    int64         `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
}
