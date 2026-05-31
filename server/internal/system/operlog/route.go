package operlog

import (
	"github.com/gin-gonic/gin"

	"server/internal/shared/database"
)

func Register(r *gin.RouterGroup) {
	repo := NewRepository(database.DB)
	svc := NewService(repo)
	handler := NewHandler(svc)

	handler.RegisterRoutes(r)
}
