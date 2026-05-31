package auth

import (
	"github.com/gin-gonic/gin"

	"server/internal/shared/database"
	"server/internal/shared/jwt"
	"server/internal/system/user"
)

// Register 注册Auth模块路由
func Register(r *gin.RouterGroup, jwtManager *jwt.JWT) {
	userRepo := user.NewRepository(database.DB)
	svc := NewService(userRepo)
	handler := NewHandler(svc, jwtManager)

	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
	r.POST("/forgot-password", handler.ForgotPassword)
	r.POST("/reset-password", handler.ResetPassword)
	r.GET("/captcha", handler.Captcha)
}
