package jobs

import (
	"go.uber.org/zap"
)

// TaskDemo 实现 Task 接口的一个示例
type TaskDemo struct {
	logger *zap.Logger
}

// NewTaskDemo
func NewTaskDemo(l *zap.Logger) *TaskDemo {
	return &TaskDemo{
		logger: l,
	}
}

func (this *TaskDemo) Run() {
	this.logger.Info("TaskDemo is running")
}
