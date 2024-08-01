package jobs

import (
	"fmt"
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

func (te *TaskDemo) Run() {
	fmt.Println("Running task...")
}
