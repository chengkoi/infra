package auth

import (
	"github.com/gin-gonic/gin"

	"server/internal/config"
	"server/internal/shared/database"
	"server/internal/shared/jwt"
	"server/internal/shared/utils"
	"server/internal/system/loginlog"
	"server/internal/system/user"
)

// Register 注册Auth模块路由
func Register(r *gin.RouterGroup, jwtManager *jwt.JWT) {
	userRepo := user.NewRepository(database.DB)
	loginlogRepo := loginlog.NewRepository(database.DB)
	svc := NewService(userRepo)

	captchaCfg := utils.CaptchaConfig{
		Enabled: config.Conf.Captcha.Enabled,
		Mode:    config.Conf.Captcha.Mode,
		Length:  config.Conf.Captcha.Length,
		Width:   config.Conf.Captcha.Width,
		Height:  config.Conf.Captcha.Height,
		Expire:  config.Conf.Captcha.Expire,
	}

	handler := NewHandler(svc, jwtManager, loginlogRepo, captchaCfg)

	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
	r.POST("/forgot-password", handler.ForgotPassword)
	r.POST("/reset-password", handler.ResetPassword)
	r.GET("/captcha", handler.Captcha)
}
