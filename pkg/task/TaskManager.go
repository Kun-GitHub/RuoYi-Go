package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
	"time"
)

// Task 接口定义了一个任务必须实现的方法
type Task interface {
	Run()
}

// TaskManager 用于管理多个任务
type TaskManager struct {
	mu         sync.Mutex
	tasks      map[string]Task
	schedulers map[string]*cron.Cron
}

// NewTaskManager 创建一个新的任务管理器实例
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:      make(map[string]Task),
		schedulers: make(map[string]*cron.Cron),
	}
}

// RegisterTask 注册一个任务
func (tm *TaskManager) RegisterTask(name string, task Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.tasks[name] = task
}

// StartTask 启动一个基于Cron表达式的时间任务
func (tm *TaskManager) StartTask(name string, cronSpec string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, ok := tm.tasks[name]
	if !ok {
		log.Fatalf("Task not registered: %s", name)
	}

	c := cron.New()
	_, err := c.AddFunc(cronSpec, task.Run)
	if err != nil {
		log.Fatal(err)
	}

	tm.schedulers[name] = c
	c.Start()
}

// StopTask 停止一个任务
func (tm *TaskManager) StopTask(name string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	scheduler, ok := tm.schedulers[name]
	if !ok {
		log.Fatalf("Scheduler not found for task: %s", name)
	}

	scheduler.Stop()
	delete(tm.schedulers, name)
}

// GetTasks 返回已注册的任务列表
func (tm *TaskManager) GetTasks() []string {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	var names []string
	for name := range tm.tasks {
		names = append(names, name)
	}
	return names
}

// TaskExample 实现 Task 接口的一个示例
type TaskExample struct{}

func (te *TaskExample) Run() {
	fmt.Println("Running task...")
}

func main() {
	taskManager := NewTaskManager()

	// 注册一个任务
	taskManager.RegisterTask("TaskExample", &TaskExample{})

	// 启动一个基于Cron表达式的时间任务
	taskManager.StartTask("TaskExample", "0/15 * * * * ?") // 每15分钟执行一次

	// 获取所有已注册的任务名称
	tasks := taskManager.GetTasks()
	fmt.Println("Registered tasks:", tasks)

	// 运行一段时间后停止任务
	time.Sleep(2 * time.Minute)
	taskManager.StopTask("TaskExample")
}
