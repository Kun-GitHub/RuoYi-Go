package handler

import (
	"RuoYi-Go/internal/common"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"go.uber.org/zap"
)

type MonitorHandler struct {
	logger *zap.Logger
}

func NewMonitorHandler(logger *zap.Logger) *MonitorHandler {
	return &MonitorHandler{logger: logger}
}

func (h *MonitorHandler) Server(ctx iris.Context) {
	// 获取 CPU 使用率
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println("Error getting CPU percent:", err)
		return
	}
	fmt.Printf("CPU Usage: %v\n", cpuPercent)

	// 获取内存信息
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting memory stats:", err)
		return
	}
	fmt.Printf("Total Memory: %v\n", vmStat.Total)
	fmt.Printf("Free Memory: %v\n", vmStat.Free)
	fmt.Printf("Used Memory: %v\n", vmStat.Used)

	ctx.JSON(common.Success(nil))
}

func (h *MonitorHandler) Cache(ctx iris.Context) {
	ctx.JSON(common.Success(nil))
}

func (h *MonitorHandler) CacheNames(ctx iris.Context) {
	ctx.JSON(common.Success(nil))
}
