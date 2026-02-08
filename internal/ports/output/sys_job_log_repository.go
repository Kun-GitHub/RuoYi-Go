package output

import "RuoYi-Go/internal/domain/model"

type SysJobLogRepository interface {
	List(page, size int, jobName, jobGroup, status string, createTime []string) ([]*model.SysJobLog, int64, error)
	Get(id int64) (*model.SysJobLog, error)
	Delete(ids []int64) error
	Clean() error
	Create(jobLog *model.SysJobLog) error
}
