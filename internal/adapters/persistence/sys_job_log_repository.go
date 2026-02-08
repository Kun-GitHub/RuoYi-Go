package persistence

import (
	"RuoYi-Go/internal/adapters/dao"
	"RuoYi-Go/internal/domain/model"
	"context"
	"time"
)

type SysJobLogRepository struct {
	db *dao.DatabaseStruct
}

func NewSysJobLogRepository(db *dao.DatabaseStruct) *SysJobLogRepository {
	return &SysJobLogRepository{db: db}
}

func (this *SysJobLogRepository) List(page, size int, jobName, jobGroup, status string, createTime []string) ([]*model.SysJobLog, int64, error) {
	q := this.db.Gen.SysJobLog.WithContext(context.Background())

	if jobName != "" {
		q = q.Where(this.db.Gen.SysJobLog.JobName.Like("%" + jobName + "%"))
	}
	if jobGroup != "" {
		q = q.Where(this.db.Gen.SysJobLog.JobGroup.Eq(jobGroup))
	}
	if status != "" {
		q = q.Where(this.db.Gen.SysJobLog.Status.Eq(status))
	}
	if len(createTime) == 2 {
		t1, err1 := time.Parse("2006-01-02", createTime[0])
		t2, err2 := time.Parse("2006-01-02", createTime[1])
		if err1 == nil && err2 == nil {
			startOfDay := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
			endOfDay := time.Date(t2.Year(), t2.Month(), t2.Day(), 23, 59, 59, 0, t2.Location())
			q = q.Where(this.db.Gen.SysJobLog.CreateTime.Between(startOfDay, endOfDay))
		}
	}

	return q.Order(this.db.Gen.SysJobLog.JobLogID.Desc()).FindByPage((page-1)*size, size)
}

func (this *SysJobLogRepository) Get(id int64) (*model.SysJobLog, error) {
	return this.db.Gen.SysJobLog.WithContext(context.Background()).Where(this.db.Gen.SysJobLog.JobLogID.Eq(id)).First()
}

func (this *SysJobLogRepository) Delete(ids []int64) error {
	_, err := this.db.Gen.SysJobLog.WithContext(context.Background()).Where(this.db.Gen.SysJobLog.JobLogID.In(ids...)).Delete()
	return err
}

func (this *SysJobLogRepository) Clean() error {
	// Truncate or Delete all
	_, err := this.db.Gen.SysJobLog.WithContext(context.Background()).Where(this.db.Gen.SysJobLog.JobLogID.Gt(0)).Delete()
	return err
}

func (this *SysJobLogRepository) Create(jobLog *model.SysJobLog) error {
	return this.db.Gen.SysJobLog.WithContext(context.Background()).Create(jobLog)
}
