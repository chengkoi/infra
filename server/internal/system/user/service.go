package user

import (
	"errors"

	"server/internal/shared/utils"
	"gorm.io/gorm"
)

// Service 用户业务层
type Service struct {
	repo *Repository
}

// NewService 创建Service
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// CreateUser 创建用户
func (s *Service) CreateUser(req *CreateUserReq) error {
	// 检查用户名是否已存在
	exists, err := s.repo.CheckUsernameExists(req.Username, 0)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("用户名已存在")
	}

	user := &User{
		Username: req.Username,
		Password: utils.MD5(req.Password),
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Gender:   req.Gender,
		Remark:   req.Remark,
		Status:   1,
	}

	return s.repo.Create(user)
}

// UpdateUser 更新用户
func (s *Service) UpdateUser(id uint, req *UpdateUserReq) error {
	// 检查用户是否存在
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	updates := make(map[string]interface{})

	if req.Nickname != "" {
		updates["nickname"] = req.Nickname
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	updates["gender"] = req.Gender
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}

	return s.repo.Update(id, updates)
}

// DeleteUser 删除用户
func (s *Service) DeleteUser(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	return s.repo.Delete(id)
}

// GetUser 获取用户详情
func (s *Service) GetUser(id uint) (*UserResp, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	return toUserResp(user), nil
}

// ListUsers 用户列表
func (s *Service) ListUsers(query *UserQueryReq) (*UserListResp, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	users, total, err := s.repo.List(query)
	if err != nil {
		return nil, err
	}

	resp := make([]UserResp, 0, len(users))
	for _, u := range users {
		resp = append(resp, *toUserResp(&u))
	}

	return &UserListResp{
		List:     resp,
		Total:    total,
		Page:     query.Page,
		PageSize: query.PageSize,
	}, nil
}

// toUserResp 转换为响应对象
func toUserResp(u *User) *UserResp {
	return &UserResp{
		ID:        u.ID,
		Username:  u.Username,
		Nickname:  u.Nickname,
		Email:     u.Email,
		Phone:     u.Phone,
		Avatar:    u.Avatar,
		Gender:    u.Gender,
		Status:    u.Status,
		Remark:    u.Remark,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}