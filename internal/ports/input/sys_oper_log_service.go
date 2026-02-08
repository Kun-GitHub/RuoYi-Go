package input

import "RuoYi-Go/internal/domain/model"

type SysOperLogService interface {
	List(page, size int, operIp, title, operName, businessType, status string, operTime []string) ([]*model.SysOperLog, int64, error)
	Get(id int64) (*model.SysOperLog, error)
	Delete(ids string) error
	Clean() error
	Create(operLog *model.SysOperLog) error
}
