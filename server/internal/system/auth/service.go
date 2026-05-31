package auth

import (
	"errors"

	"server/internal/shared/utils"
	"server/internal/system/user"
	"gorm.io/gorm"
)

// Service 认证业务层
type Service struct {
	userRepo *user.Repository
}

// NewService 创建Service
func NewService(userRepo *user.Repository) *Service {
	return &Service{userRepo: userRepo}
}

// Login 用户登录，返回用户实体供 handler 层生成 token
func (s *Service) Login(req *LoginReq) (*user.User, error) {
	u, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}

	if u.Status == 0 {
		return nil, errors.New("用户已被禁用")
	}

	if utils.MD5(req.Password) != u.Password {
		return nil, errors.New("用户名或密码错误")
	}

	return u, nil
}

// Register 用户注册
func (s *Service) Register(req *RegisterReq) error {
	exists, err := s.userRepo.CheckUsernameExists(req.Username, 0)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("用户名已存在")
	}

	u := &user.User{
		Username: req.Username,
		Password: utils.MD5(req.Password),
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Status:   1,
	}

	return s.userRepo.Create(u)
}

// ForgotPassword 忘记密码：返回新密码
func (s *Service) ForgotPassword(username string) (string, error) {
	u, err := s.userRepo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("用户不存在")
		}
		return "", err
	}

	if u.Status == 0 {
		return "", errors.New("用户已被禁用")
	}

	// 生成随机新密码
	newPassword := utils.RandomString(10)
	hashedPassword := utils.MD5(newPassword)

	if err := s.userRepo.Update(u.ID, map[string]interface{}{"password": hashedPassword}); err != nil {
		return "", err
	}

	return newPassword, nil
}

// ResetPassword 重置密码：旧密码校验 + 更新为新密码
func (s *Service) ResetPassword(req *ResetPasswordReq) error {
	u, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	// 校验旧密码
	if utils.MD5(req.OldPassword) != u.Password {
		return errors.New("旧密码错误")
	}

	// 更新密码
	return s.userRepo.Update(u.ID, map[string]interface{}{"password": utils.MD5(req.NewPassword)})
}
