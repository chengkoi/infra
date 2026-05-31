package auth

// LoginReq 登录请求
type LoginReq struct {
	Username   string `json:"username" binding:"required" validate:"required,min=3,max=64"`
	Password   string `json:"password" binding:"required" validate:"required,min=6,max=32"`
	CaptchaID  string `json:"captcha_id" binding:"required"`
	CaptchaAns string `json:"captcha_ans" binding:"required"`
}

// RegisterReq 注册请求
type RegisterReq struct {
	Username   string `json:"username" binding:"required" validate:"required,min=3,max=64"`
	Password   string `json:"password" binding:"required" validate:"required,min=6,max=32"`
	Nickname   string `json:"nickname" binding:"required" validate:"required,max=64"`
	Email      string `json:"email" validate:"omitempty,email,max=128"`
	Phone      string `json:"phone" validate:"omitempty,max=20"`
	CaptchaID  string `json:"captcha_id" binding:"required"`
	CaptchaAns string `json:"captcha_ans" binding:"required"`
}

// ForgotPasswordReq 忘记密码请求
type ForgotPasswordReq struct {
	Username   string `json:"username" binding:"required"`
	CaptchaID  string `json:"captcha_id" binding:"required"`
	CaptchaAns string `json:"captcha_ans" binding:"required"`
}

// ResetPasswordReq 重置密码请求
type ResetPasswordReq struct {
	Username    string `json:"username" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required" validate:"required,min=6,max=32"`
}

// LoginResp 登录响应
type LoginResp struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expire_at"`
}
