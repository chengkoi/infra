package operlog

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(log *OperLog) error {
	return r.db.Create(log).Error
}

func (r *Repository) Delete(ids []uint) error {
	return r.db.Delete(&OperLog{}, ids).Error
}

func (r *Repository) List(query *OperLogQueryReq) ([]OperLog, int64, error) {
	var logs []OperLog
	var total int64

	db := r.db.Model(&OperLog{})
	if query.Username != "" {
		db = db.Where("username LIKE ?", "%"+query.Username+"%")
	}
	if query.Module != "" {
		db = db.Where("module = ?", query.Module)
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
	return r.db.Where("1 = 1").Delete(&OperLog{}).Error
}
