package jobs

import (
	"go.uber.org/zap"
)

// TaskDemo 实现 Task 接口的一个示例
type TaskGoroutine struct {
	logger      *zap.Logger
	pageNumChan chan int
}

// NewTaskDemo
func NewTaskGoroutine(l *zap.Logger) *TaskGoroutine {
	pageNumChan := make(chan int, 1)
	pageNumChan <- 0 // 初始化pageNum为0

	return &TaskGoroutine{
		logger:      l,
		pageNumChan: pageNumChan,
	}
}

func (this *TaskGoroutine) Run() {
	this.logger.Info("TaskGoroutine is running")
	select {
	case pageNum := <-this.pageNumChan:
		this.logger.Info("当前页码为：%d", zap.Int("pageNum", pageNum))
		this.fetchAndInsertData(pageNum)
	default:
		defaultPageNum := 0
		this.fetchAndInsertData(defaultPageNum)
	}
}

func (this *TaskGoroutine) fetchAndInsertData(pageNum int) {
	this.pageNumChan <- pageNum + 1
	// 获取数据
	// 将数据插入数据库
}
