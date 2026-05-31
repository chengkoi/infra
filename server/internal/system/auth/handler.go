package auth

import (
	"github.com/gin-gonic/gin"

	"server/internal/shared/errno"
	"server/internal/shared/jwt"
	"server/internal/shared/response"
	"server/internal/shared/utils"
)

// Handler 认证HTTP处理器
type Handler struct {
	svc        *Service
	jwtManager *jwt.JWT
}

// NewHandler 创建Handler
func NewHandler(svc *Service, jwtManager *jwt.JWT) *Handler {
	return &Handler{svc: svc, jwtManager: jwtManager}
}

// Login 用户登录
// @Summary      用户登录
// @Description  用户名密码 + 验证码登录
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      LoginReq  true  "登录请求"
// @Success      200   {object}  response.Response{data=LoginResp}
// @Failure      400   {object}  response.Response
// @Router       /api/v1/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	if !utils.VerifyCaptcha(req.CaptchaID, req.CaptchaAns) {
		response.Fail(c, errno.CaptchaError, "验证码错误或已过期")
		return
	}

	u, err := h.svc.Login(&req)
	if err != nil {
		response.Fail(c, errno.PasswordError, err.Error())
		return
	}

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
// @Summary      用户注册
// @Description  用户注册
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      RegisterReq  true  "注册请求"
// @Success      200   {object}  response.Response
// @Failure      400   {object}  response.Response
// @Router       /api/v1/register [post]
func (h *Handler) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	if !utils.VerifyCaptcha(req.CaptchaID, req.CaptchaAns) {
		response.Fail(c, errno.CaptchaError, "验证码错误或已过期")
		return
	}

	if err := h.svc.Register(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	response.SuccessWithMessage(c, "注册成功", nil)
}

// ForgotPassword 忘记密码
// @Summary      忘记密码
// @Description  通过用户名 + 验证码获取新密码
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      ForgotPasswordReq  true  "忘记密码请求"
// @Success      200   {object}  response.Response
// @Failure      400   {object}  response.Response
// @Router       /api/v1/forgot-password [post]
func (h *Handler) ForgotPassword(c *gin.Context) {
	var req ForgotPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	if !utils.VerifyCaptcha(req.CaptchaID, req.CaptchaAns) {
		response.Fail(c, errno.CaptchaError, "验证码错误或已过期")
		return
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
// @Summary      重置密码
// @Description  旧密码校验后设置新密码
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      ResetPasswordReq  true  "重置密码请求"
// @Success      200   {object}  response.Response
// @Failure      400   {object}  response.Response
// @Router       /api/v1/reset-password [post]
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
// @Summary      获取验证码
// @Description  获取数字验证码
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response{data=utils.CaptchaResult}
// @Router       /api/v1/captcha [get]
func (h *Handler) Captcha(c *gin.Context) {
	result, err := utils.GenerateCaptcha()
	if err != nil {
		response.Fail(c, errno.InternalError, "生成验证码失败")
		return
	}
	response.Success(c, result)
}
