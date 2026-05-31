package errno

const (
	Success         = 0
	BadRequest      = 40000
	Unauthorized    = 40100
	Forbidden       = 40300
	NotFound        = 40400
	UserNotFound    = 40001
	RoleNotFound    = 40002
	PermissionError = 40301
	InternalError   = 50000
	DatabaseError   = 50001
)

var errMsg = map[int]string{
	Success:         "success",
	BadRequest:      "请求参数错误",
	Unauthorized:    "未授权",
	Forbidden:       "禁止访问",
	NotFound:        "资源不存在",
	UserNotFound:    "用户不存在",
	RoleNotFound:    "角色不存在",
	PermissionError: "权限不足",
	InternalError:   "内部服务器错误",
	DatabaseError:   "数据库错误",
}

// Errno 错误码
type Errno struct {
	Code int
	Msg  string
}

// Error 实现error接口
func (e *Errno) Error() string {
	return e.Msg
}

// New 创建错误
func New(code int, msg string) *Errno {
	return &Errno{
		Code: code,
		Msg:  msg,
	}
}

// Decode 解码错误
func Decode(err error) (int, string) {
	if err == nil {
		return Success, errMsg[Success]
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.Code, typed.Msg
	default:
		return InternalError, errMsg[InternalError]
	}
}

// GetMessage 获取错误消息
func GetMessage(code int) string {
	if msg, ok := errMsg[code]; ok {
		return msg
	}
	return "未知错误"
}