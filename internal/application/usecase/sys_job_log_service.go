package usecase

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"

	"go.uber.org/zap"
)

type SysJobLogService struct {
	repo   output.SysJobLogRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysJobLogService(repo output.SysJobLogRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysJobLogService {
	return &SysJobLogService{repo: repo, cache: cache, logger: logger}
}

func (s *SysJobLogService) List(page, size int, jobName, jobGroup, status string, createTime []string) ([]*model.SysJobLog, int64, error) {
	return s.repo.List(page, size, jobName, jobGroup, status, createTime)
}

func (s *SysJobLogService) Get(id int64) (*model.SysJobLog, error) {
	return s.repo.Get(id)
}

func (s *SysJobLogService) Delete(ids string) error {
	idList := common.SplitInt64(ids)
	return s.repo.Delete(idList)
}

func (s *SysJobLogService) Clean() error {
	return s.repo.Clean()
}

func (s *SysJobLogService) Create(jobLog *model.SysJobLog) error {
	return s.repo.Create(jobLog)
}
