// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"sync"
)

// TaskManager 用于管理多个任务
type TaskManager struct {
	mu     sync.Mutex
	logger *zap.Logger

	tasks      map[string]cron.Job
	schedulers map[string]*cron.Cron
}

// NewTaskManager 创建一个新的任务管理器实例
func NewTaskManager(l *zap.Logger) *TaskManager {
	return &TaskManager{
		logger:     l,
		tasks:      make(map[string]cron.Job),
		schedulers: make(map[string]*cron.Cron),
	}
}

// RegisterTask 注册一个任务
func (tm *TaskManager) RegisterTask(name string, task cron.Job) {
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
		tm.logger.Error(fmt.Sprintf("Task not registered: %s", name))
		return
	}

	// 创建一个新的cron.Cron实例
	scheduler, err := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow).Parse(cronSpec)
	if err != nil {
		tm.logger.Error("Failed to parse Cron expression: %s", zap.Error(err))
		return
	}

	c := cron.New()
	entryID := c.Schedule(scheduler, task)
	if entryID == 0 {
		tm.logger.Error(fmt.Sprintf("Task Fatal AddFunc: %s, err:%v", name, zap.Error(err)))
		return
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
		tm.logger.Error(fmt.Sprintf("Scheduler not found for task: %s", name))
		return
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
