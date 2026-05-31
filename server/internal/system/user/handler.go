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
// @Summary      创建用户
// @Description  创建用户
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        body  body      CreateUserReq  true  "创建用户请求"
// @Success      200   {object}  response.Response{data=UserResp}
// @Failure      400   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /api/v1/users [post]
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
// @Summary      更新用户
// @Description  更新用户
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id    path      int            true  "用户ID"
// @Param        body  body      UpdateUserReq  true  "更新用户请求"
// @Success      200   {object}  response.Response
// @Failure      400   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /api/v1/users/{id} [put]
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
// @Summary      删除用户
// @Description  删除用户
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用户ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/users/{id} [delete]
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
// @Summary      获取用户详情
// @Description  获取用户详情
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用户ID"
// @Success      200  {object}  response.Response{data=UserResp}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /api/v1/users/{id} [get]
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
// @Summary      用户列表
// @Description  查询用户列表
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        page       query     int     false  "页码"
// @Param        pageSize   query     int     false  "每页数量"
// @Param        username   query     string  false  "用户名"
// @Param        nickname   query     string  false  "昵称"
// @Param        status     query     int     false  "状态"
// @Success      200        {object}  response.Response{data=UserListResp}
// @Failure      500        {object}  response.Response
// @Router       /api/v1/users [get]
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