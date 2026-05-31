package auth

import (
	"github.com/gin-gonic/gin"

	"server/internal/shared/constants"
	"server/internal/shared/errno"
	"server/internal/shared/jwt"
	"server/internal/shared/response"
	"server/internal/shared/utils"
	"server/internal/system/loginlog"
)

// Handler 认证HTTP处理器
type Handler struct {
	svc          *Service
	jwtManager   *jwt.JWT
	loginlogRepo *loginlog.Repository
	captchaCfg   utils.CaptchaConfig
}

// NewHandler 创建Handler
func NewHandler(svc *Service, jwtManager *jwt.JWT, loginlogRepo *loginlog.Repository, captchaCfg utils.CaptchaConfig) *Handler {
	return &Handler{svc: svc, jwtManager: jwtManager, loginlogRepo: loginlogRepo, captchaCfg: captchaCfg}
}

// Login 用户登录
func (h *Handler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	// 验证码校验（启用时才校验）
	if h.captchaCfg.Enabled {
		if !utils.VerifyCaptcha(req.CaptchaID, req.CaptchaAns) {
			response.Fail(c, errno.CaptchaError, "验证码错误或已过期")
			return
		}
	}

	u, err := h.svc.Login(&req)
	if err != nil {
		loginlog.Write(h.loginlogRepo, &loginlog.LoginLog{
			Username:  req.Username,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Status:    constants.LogStatusFail,
			Message:   err.Error(),
		})
		response.Fail(c, errno.PasswordError, err.Error())
		return
	}

	loginlog.Write(h.loginlogRepo, &loginlog.LoginLog{
		UserID:    u.ID,
		Username:  u.Username,
		IP:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		Status:    constants.LogStatusSuccess,
		Message:   "登录成功",
	})

	expire := h.jwtManager.GetExpire()
	token, err := h.jwtManager.GenerateToken(u.ID, u.Username, 0, expire)
	if err != nil {
		response.Fail(c, errno.InternalError, "生成token失败")
		return
	}

	response.Success(c, &LoginResp{
		Token:    token,
		ExpireAt: expire,
	})
}

// Register 用户注册
func (h *Handler) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	if h.captchaCfg.Enabled {
		if !utils.VerifyCaptcha(req.CaptchaID, req.CaptchaAns) {
			response.Fail(c, errno.CaptchaError, "验证码错误或已过期")
			return
		}
	}

	if err := h.svc.Register(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	response.SuccessWithMessage(c, "注册成功", nil)
}

// ForgotPassword 忘记密码
func (h *Handler) ForgotPassword(c *gin.Context) {
	var req ForgotPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	if h.captchaCfg.Enabled {
		if !utils.VerifyCaptcha(req.CaptchaID, req.CaptchaAns) {
			response.Fail(c, errno.CaptchaError, "验证码错误或已过期")
			return
		}
	}

	newPassword, err := h.svc.ForgotPassword(req.Username)
	if err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	response.SuccessWithMessage(c, "新密码已生成，请登录后修改", gin.H{
		"new_password": newPassword,
	})
}

// ResetPassword 重置密码
func (h *Handler) ResetPassword(c *gin.Context) {
	var req ResetPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	if err := h.svc.ResetPassword(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	response.SuccessWithMessage(c, "密码重置成功", nil)
}

// Captcha 获取验证码
func (h *Handler) Captcha(c *gin.Context) {
	// 如果未启用验证码，返回空数据表示跳过
	if !h.captchaCfg.Enabled {
		response.Success(c, map[string]any{
			"captcha_enabled": false,
		})
		return
	}

	result, err := utils.GenerateCaptcha(h.captchaCfg)
	if err != nil {
		response.Fail(c, errno.InternalError, "生成验证码失败")
		return
	}
	result.CaptchaID = result.CaptchaID
	response.Success(c, result)
}
