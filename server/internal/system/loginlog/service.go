package loginlog

import "time"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(query *LoginLogQueryReq) (*LoginLogListResp, error) {
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

	resp := make([]LoginLogResp, 0, len(logs))
	for _, l := range logs {
		resp = append(resp, LoginLogResp{
			ID:        l.ID,
			UserID:    l.UserID,
			Username:  l.Username,
			IP:        l.IP,
			UserAgent: l.UserAgent,
			Status:    l.Status,
			Message:   l.Message,
			CreatedAt: l.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &LoginLogListResp{List: resp, Total: total, Page: query.Page, PageSize: query.PageSize}, nil
}

func (s *Service) Delete(ids []uint) error {
	return s.repo.Delete(ids)
}

func (s *Service) Clear() error {
	return s.repo.Clear()
}

// Write 写入登录日志（由auth handler调用）
func Write(repo *Repository, log *LoginLog) {
	log.CreatedAt = time.Now()
	_ = repo.Create(log)
}
