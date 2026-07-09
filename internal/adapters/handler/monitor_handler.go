package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/pkg/cache"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"go.uber.org/zap"
)

type MonitorHandler struct {
	logger *zap.Logger
	redis  *cache.RedisClient
}

func NewMonitorHandler(logger *zap.Logger, redis *cache.RedisClient) *MonitorHandler {
	return &MonitorHandler{logger: logger, redis: redis}
}

func (h *MonitorHandler) Server(ctx iris.Context) {
	var result model.MonitorServer

	// CPU
	result.Cpu = collectCpuInfo()

	// Memory
	result.Mem = collectMemInfo()

	// JVM (Go runtime equivalent)
	result.Jvm = collectJvmInfo()

	// System / OS
	result.Sys = collectSysInfo()

	// Disk
	result.SysFiles = collectSysFiles()

	ctx.JSON(common.Success(result))
}

func collectCpuInfo() *model.CpuInfo {
	info := &model.CpuInfo{}

	cpuCount, _ := cpu.Counts(true)
	info.CpuNum = cpuCount

	// CPU time stats (total ticks across all CPUs)
	times, err := cpu.Times(false)
	if err != nil || len(times) == 0 {
		return info
	}
	t := times[0]
	total := t.User + t.System + t.Idle + t.Iowait + t.Irq + t.Softirq + t.Steal
	info.Total = total
	info.Sys = t.System
	info.Used = t.User
	info.Wait = t.Iowait
	info.Free = t.Idle

	return info
}

func collectMemInfo() *model.MemInfo {
	info := &model.MemInfo{}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return info
	}

	info.Total = float64(vm.Total) / 1024 / 1024 / 1024 // GB
	info.Used = float64(vm.Used) / 1024 / 1024 / 1024
	info.Free = float64(vm.Available) / 1024 / 1024 / 1024
	info.Usage = vm.UsedPercent

	return info
}

func collectJvmInfo() *model.JvmInfo {
	info := &model.JvmInfo{}

	// Go runtime stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	totalMB := float64(m.Sys) / 1024 / 1024
	usedMB := float64(m.HeapAlloc) / 1024 / 1024
	freeMB := totalMB - usedMB
	if freeMB < 0 {
		freeMB = 0
	}
	maxMB := totalMB
	usagePct := 0.0
	if totalMB > 0 {
		usagePct = usedMB / totalMB * 100
	}

	info.Total = totalMB
	info.Max = maxMB
	info.Free = freeMB
	info.Used = usedMB
	info.Usage = usagePct

	// Runtime name
	info.Name = "Go Runtime"

	// Go version
	info.Version = strings.TrimPrefix(runtime.Version(), "go")

	// GOROOT
	info.Home = runtime.GOROOT()

	// Process start time
	info.StartTime = "N/A"
	info.RunTime = fmt.Sprintf("%d hours", int(time.Since(processStartTime).Hours()))

	// Args
	info.InputArgs = strings.Join(os.Args[1:], " ")

	return info
}

var processStartTime = time.Now()

func collectSysInfo() *model.SysInfo {
	info := &model.SysInfo{}

	hostInfo, err := host.Info()
	if err == nil {
		info.ComputerName = hostInfo.Hostname
		info.OsName = hostInfo.Platform + " " + hostInfo.PlatformVersion
	}

	info.ComputerIp = getLocalIp()

	dir, _ := os.Getwd()
	info.UserDir = dir

	info.OsArch = runtime.GOARCH

	return info
}

func collectSysFiles() []*model.SysFile {
	var files []*model.SysFile

	partitions, err := disk.Partitions(false)
	if err != nil {
		return files
	}

	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}

		usagePct := 0.0
		if usage.Total > 0 {
			usagePct = usage.UsedPercent
		}

		files = append(files, &model.SysFile{
			DirName:     p.Mountpoint,
			SysTypeName: p.Fstype,
			TypeName:    p.Device,
			Total:       formatBytes(usage.Total),
			Free:        formatBytes(usage.Free),
			Used:        formatBytes(usage.Used),
			Usage:       usagePct,
		})
	}

	return files
}

func formatBytes(b uint64) string {
	const unit = 1024
	const unitName = "GB"
	value := float64(b) / unit / unit / unit
	return fmt.Sprintf("%.1f %s", value, unitName)
}

func getLocalIp() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, a := range addr {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

func (h *MonitorHandler) Cache(ctx iris.Context) {
	ctx.JSON(common.Success(nil))
}

func (h *MonitorHandler) CacheNames(ctx iris.Context) {
	ctx.JSON(common.Success(nil))
}

func (h *MonitorHandler) GetKeys(ctx iris.Context) {
	cacheName := ctx.Params().GetString("cacheName")
	keys, _ := h.redis.Keys(fmt.Sprintf("%s*", cacheName))
	ctx.JSON(common.Success(keys))
}

func (h *MonitorHandler) GetValue(ctx iris.Context) {
	cacheName := ctx.Params().GetString("cacheName")
	cacheKey := ctx.Params().GetString("cacheKey")
	val, _ := h.redis.Get(fmt.Sprintf("%s%s", cacheName, cacheKey))
	ctx.JSON(common.Success(map[string]string{"value": val}))
}

func (h *MonitorHandler) ClearCacheName(ctx iris.Context) {
	cacheName := ctx.Params().GetString("cacheName")
	keys, _ := h.redis.Keys(fmt.Sprintf("%s*", cacheName))
	for _, key := range keys {
		h.redis.Del(key)
	}
	ctx.JSON(common.Success(nil))
}

func (h *MonitorHandler) ClearCacheKey(ctx iris.Context) {
	cacheName := ctx.Params().GetString("cacheName")
	cacheKey := ctx.Params().GetString("cacheKey")
	h.redis.Del(fmt.Sprintf("%s%s", cacheName, cacheKey))
	ctx.JSON(common.Success(nil))
}

func (h *MonitorHandler) ClearCacheAll(ctx iris.Context) {
	keys, _ := h.redis.Keys("*")
	for _, key := range keys {
		h.redis.Del(key)
	}
	ctx.JSON(common.Success(nil))
}
