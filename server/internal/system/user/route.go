package user

import (
	"github.com/gin-gonic/gin"

	"server/internal/shared/database"
)

// Register 注册User模块路由
func Register(r *gin.RouterGroup) {
	repo := NewRepository(database.DB)
	svc := NewService(repo)
	handler := NewHandler(svc)

	handler.RegisterRoutes(r)
}