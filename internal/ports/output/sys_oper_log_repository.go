package output

import "RuoYi-Go/internal/domain/model"

type SysOperLogRepository interface {
	List(page, size int, operIp, title, operName, businessType, status string, operTime []string) ([]*model.SysOperLog, int64, error)
	Get(id int64) (*model.SysOperLog, error)
	Delete(ids []int64) error
	Clean() error
	Create(operLog *model.SysOperLog) error
}
