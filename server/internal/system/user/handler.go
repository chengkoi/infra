package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"server/internal/shared/errno"
	"server/internal/shared/response"
)

// Handler 用户HTTP处理器
type Handler struct {
	svc *Service
}

// NewHandler 创建Handler
func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// CreateUser 创建用户
func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	if err := h.svc.CreateUser(&req); err != nil {
		response.Fail(c, errno.InternalError, err.Error())
		return
	}

	response.SuccessWithMessage(c, "创建成功", nil)
}

// UpdateUser 更新用户
func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errno.BadRequest, "无效的ID")
		return
	}

	var req UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	if err := h.svc.UpdateUser(uint(id), &req); err != nil {
		response.Fail(c, errno.InternalError, err.Error())
		return
	}

	response.SuccessWithMessage(c, "更新成功", nil)
}

// DeleteUser 删除用户
func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errno.BadRequest, "无效的ID")
		return
	}

	if err := h.svc.DeleteUser(uint(id)); err != nil {
		response.Fail(c, errno.InternalError, err.Error())
		return
	}

	response.SuccessWithMessage(c, "删除成功", nil)
}

// GetUser 获取用户详情
func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, errno.BadRequest, "无效的ID")
		return
	}

	user, err := h.svc.GetUser(uint(id))
	if err != nil {
		response.Fail(c, errno.UserNotFound, err.Error())
		return
	}

	response.Success(c, user)
}

// ListUsers 用户列表
func (h *Handler) ListUsers(c *gin.Context) {
	var query UserQueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Fail(c, errno.BadRequest, err.Error())
		return
	}

	result, err := h.svc.ListUsers(&query)
	if err != nil {
		response.Fail(c, errno.InternalError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Code: response.SuccessCode,
		Msg:  response.SuccessMsg,
		Data: result,
	})
}

// RegisterRoutes 注册User路由
func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	group := r.Group("/users")
	{
		group.GET("", h.ListUsers)
		group.GET("/:id", h.GetUser)
		group.POST("", h.CreateUser)
		group.PUT("/:id", h.UpdateUser)
		group.DELETE("/:id", h.DeleteUser)
	}
}