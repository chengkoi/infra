package user

import (
	"gorm.io/gorm"
)

// Repository 用户数据访问层
type Repository struct {
	db *gorm.DB
}

// NewRepository 创建Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Create 创建用户
func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}

// Update 更新用户
func (r *Repository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&User{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除用户
func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

// GetByID 根据ID获取用户
func (r *Repository) GetByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername 根据用户名获取用户
func (r *Repository) GetByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// List 用户列表
func (r *Repository) List(query *UserQueryReq) ([]User, int64, error) {
	var users []User
	var total int64

	db := r.db.Model(&User{})

	if query.Username != "" {
		db = db.Where("username LIKE ?", "%"+query.Username+"%")
	}
	if query.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+query.Nickname+"%")
	}
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset((query.Page - 1) * query.PageSize).
		Limit(query.PageSize).
		Order("id DESC").
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// UpdateStatus 更新状态
func (r *Repository) UpdateStatus(id uint, status int) error {
	return r.db.Model(&User{}).Where("id = ?", id).Update("status", status).Error
}

// CheckUsernameExists 检查用户名是否存在
func (r *Repository) CheckUsernameExists(username string, excludeID uint) (bool, error) {
	var count int64
	db := r.db.Model(&User{}).Where("username = ?", username)
	if excludeID > 0 {
		db = db.Where("id != ?", excludeID)
	}
	if err := db.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}