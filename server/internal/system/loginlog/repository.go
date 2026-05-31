package loginlog

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(log *LoginLog) error {
	return r.db.Create(log).Error
}

func (r *Repository) Delete(ids []uint) error {
	return r.db.Delete(&LoginLog{}, ids).Error
}

func (r *Repository) List(query *LoginLogQueryReq) ([]LoginLog, int64, error) {
	var logs []LoginLog
	var total int64

	db := r.db.Model(&LoginLog{})
	if query.Username != "" {
		db = db.Where("username LIKE ?", "%"+query.Username+"%")
	}
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset((query.Page - 1) * query.PageSize).
		Limit(query.PageSize).Order("id DESC").
		Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

func (r *Repository) Clear() error {
	return r.db.Where("1 = 1").Delete(&LoginLog{}).Error
}
