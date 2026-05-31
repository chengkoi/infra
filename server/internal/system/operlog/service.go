package operlog

import (
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(query *OperLogQueryReq) (*OperLogListResp, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	logs, total, err := s.repo.List(query)
	if err != nil {
		return nil, err
	}

	resp := make([]OperLogResp, 0, len(logs))
	for _, l := range logs {
		resp = append(resp, OperLogResp{
			ID:        l.ID,
			UserID:    l.UserID,
			Username:  l.Username,
			Module:    l.Module,
			Action:    l.Action,
			Method:    l.Method,
			Path:      l.Path,
			IP:        l.IP,
			UserAgent: l.UserAgent,
			Duration:  l.Duration,
			Status:    l.Status,
			CreatedAt: l.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &OperLogListResp{List: resp, Total: total, Page: query.Page, PageSize: query.PageSize}, nil
}

func (s *Service) Delete(ids []uint) error {
	return s.repo.Delete(ids)
}

func (s *Service) Clear() error {
	return s.repo.Clear()
}

// Write 写入操作日志（由中间件调用）
func Write(repo *Repository, log *OperLog) {
	log.CreatedAt = time.Now()
	_ = repo.Create(log)
}
