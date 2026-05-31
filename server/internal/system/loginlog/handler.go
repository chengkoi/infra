package loginlog

import (
	"github.com/gin-gonic/gin"

	"server/internal/shared/errno"
	"server/internal/shared/response"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// List 登录日志列表
func (h *Handler) List(c *gin.Context) {
	var query LoginLogQueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	result, err := h.svc.List(&query)
	if err != nil {
		response.Fail(c, errno.InternalError, err.Error())
		return
	}

	response.Success(c, result)
}

// Delete 删除登录日志
func (h *Handler) Delete(c *gin.Context) {
	var ids []uint
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	if err := h.svc.Delete(ids); err != nil {
		response.Fail(c, errno.InternalError, err.Error())
		return
	}

	response.SuccessWithMessage(c, "删除成功", nil)
}

// Clear 清空登录日志
func (h *Handler) Clear(c *gin.Context) {
	if err := h.svc.Clear(); err != nil {
		response.Fail(c, errno.InternalError, err.Error())
		return
	}

	response.SuccessWithMessage(c, "清空成功", nil)
}

// RegisterRoutes 注册路由
func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	group := r.Group("/login-logs")
	{
		group.GET("", h.List)
		group.DELETE("", h.Delete)
		group.DELETE("/clear", h.Clear)
	}
}
