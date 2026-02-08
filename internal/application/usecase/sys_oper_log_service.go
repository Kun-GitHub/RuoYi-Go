package usecase

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"

	"go.uber.org/zap"
)

type SysOperLogService struct {
	repo   output.SysOperLogRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysOperLogService(repo output.SysOperLogRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysOperLogService {
	return &SysOperLogService{repo: repo, cache: cache, logger: logger}
}

func (s *SysOperLogService) List(page, size int, operIp, title, operName, businessType, status string, operTime []string) ([]*model.SysOperLog, int64, error) {
	return s.repo.List(page, size, operIp, title, operName, businessType, status, operTime)
}

func (s *SysOperLogService) Get(id int64) (*model.SysOperLog, error) {
	return s.repo.Get(id)
}

func (s *SysOperLogService) Delete(ids string) error {
	idList := common.SplitInt64(ids)
	return s.repo.Delete(idList)
}

func (s *SysOperLogService) Clean() error {
	return s.repo.Clean()
}

func (s *SysOperLogService) Create(operLog *model.SysOperLog) error {
	return s.repo.Create(operLog)
}
