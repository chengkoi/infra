package constants

const (
	// Status 状态
	StatusEnable  = 1
	StatusDisable = 0

	// Gender 性别
	GenderUnknown = 0
	GenderMale    = 1
	GenderFemale  = 2

	// ContextKey 上下文键
	ContextKeyUserID   = "user_id"
	ContextKeyUsername = "username"
	ContextKeyRoleID   = "role_id"
)

const (
	// LogType 日志类型
	LogTypeOperLog = 1 // 操作日志
	LogTypeLogin   = 2 // 登录日志

	// LogStatus 日志状态
	LogStatusSuccess = 1
	LogStatusFail    = 0
)

const (
	// RedisKey Redis键前缀
	RedisKeyToken  = "token:"
	RedisKeyCache  = "cache:"
	RedisKeyOnline = "online:"
)