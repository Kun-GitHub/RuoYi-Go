package persistence

import (
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/domain/model"
	"context"
	"strconv"
	"time"
)

type SysOperLogRepository struct {
	db *dao.DatabaseStruct
}

func NewSysOperLogRepository(db *dao.DatabaseStruct) *SysOperLogRepository {
	return &SysOperLogRepository{db: db}
}

func (this *SysOperLogRepository) List(page, size int, operIp, title, operName, businessType, status string, operTime []string) ([]*model.SysOperLog, int64, error) {
	q := this.db.Gen.SysOperLog.WithContext(context.Background())

	if operIp != "" {
		q = q.Where(this.db.Gen.SysOperLog.OperIP.Like("%" + operIp + "%"))
	}
	if title != "" {
		q = q.Where(this.db.Gen.SysOperLog.Title.Like("%" + title + "%"))
	}
	if operName != "" {
		q = q.Where(this.db.Gen.SysOperLog.OperName.Like("%" + operName + "%"))
	}
	if businessType != "" {
		bt, _ := strconv.Atoi(businessType)
		q = q.Where(this.db.Gen.SysOperLog.BusinessType.Eq(int32(bt)))
	}
	if status != "" {
		s, _ := strconv.Atoi(status)
		q = q.Where(this.db.Gen.SysOperLog.Status.Eq(int32(s)))
	}
	if len(operTime) == 2 {
		t1, err1 := time.Parse("2006-01-02", operTime[0])
		t2, err2 := time.Parse("2006-01-02", operTime[1])
		if err1 == nil && err2 == nil {
			startOfDay := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
			endOfDay := time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
			q = q.Where(this.db.Gen.SysOperLog.OperTime.Between(startOfDay, endOfDay))
		}
	}

	return q.Order(this.db.Gen.SysOperLog.OperID.Desc()).FindByPage((page-1)*size, size)
}

func (this *SysOperLogRepository) Get(id int64) (*model.SysOperLog, error) {
	return this.db.Gen.SysOperLog.WithContext(context.Background()).Where(this.db.Gen.SysOperLog.OperID.Eq(id)).First()
}

func (this *SysOperLogRepository) Delete(ids []int64) error {
	_, err := this.db.Gen.SysOperLog.WithContext(context.Background()).Where(this.db.Gen.SysOperLog.OperID.In(ids...)).Delete()
	return err
}

func (this *SysOperLogRepository) Clean() error {
	// Truncate or Delete all
	_, err := this.db.Gen.SysOperLog.WithContext(context.Background()).Where(this.db.Gen.SysOperLog.OperID.Gt(0)).Delete()
	return err
}

func (this *SysOperLogRepository) Create(operLog *model.SysOperLog) error {
	return this.db.Gen.SysOperLog.WithContext(context.Background()).Create(operLog)
}
