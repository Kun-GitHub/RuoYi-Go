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

// ===================== Server Monitor =====================

func (h *MonitorHandler) Server(ctx iris.Context) {
	result := model.MonitorServer{
		Cpu:      collectCpuInfo(),
		Mem:      collectMemInfo(),
		Jvm:      collectJvmInfo(),
		Sys:      collectSysInfo(),
		SysFiles: collectSysFiles(),
	}
	ctx.JSON(common.Success(result))
}

func collectCpuInfo() *model.CpuInfo {
	info := &model.CpuInfo{}

	cpuCount, _ := cpu.Counts(true)
	info.CpuNum = cpuCount

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

	info.Total = float64(vm.Total) / 1024 / 1024 / 1024
	info.Used = float64(vm.Used) / 1024 / 1024 / 1024
	info.Free = float64(vm.Available) / 1024 / 1024 / 1024
	info.Usage = vm.UsedPercent

	return info
}

func collectJvmInfo() *model.JvmInfo {
	info := &model.JvmInfo{}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	totalMB := float64(m.Sys) / 1024 / 1024
	usedMB := float64(m.HeapAlloc) / 1024 / 1024
	freeMB := totalMB - usedMB
	if freeMB < 0 {
		freeMB = 0
	}
	usagePct := 0.0
	if totalMB > 0 {
		usagePct = usedMB / totalMB * 100
	}

	info.Total = totalMB
	info.Max = totalMB
	info.Free = freeMB
	info.Used = usedMB
	info.Usage = usagePct
	info.Name = "Go Runtime"
	info.Version = strings.TrimPrefix(runtime.Version(), "go")
	info.Home = runtime.GOROOT()
	info.StartTime = processStartTime.Format("2006-01-02 15:04:05")
	info.RunTime = fmt.Sprintf("%d hours %d minutes",
		int(time.Since(processStartTime).Hours()),
		int(time.Since(processStartTime).Minutes())%60)
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
	return fmt.Sprintf("%.1f GB", float64(b)/1024/1024/1024)
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

// ===================== Cache Monitor =====================

var cacheNames = []model.SysCache{
	{CacheName: "login_tokens:", Remark: "用户信息"},
	{CacheName: "sys_config:", Remark: "配置信息"},
	{CacheName: "sys_dict:", Remark: "数据字典"},
	{CacheName: "captcha_codes:", Remark: "验证码"},
	{CacheName: "repeat_submit:", Remark: "防重提交"},
	{CacheName: "rate_limit:", Remark: "限流处理"},
	{CacheName: "pwd_err_cnt:", Remark: "密码错误次数"},
}

func (h *MonitorHandler) Cache(ctx iris.Context) {
	infoRaw, err := h.redis.Do("INFO")
	if err != nil {
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "Redis INFO failed: %s", err.Error()))
		return
	}

	cmdRaw, _ := h.redis.Do("INFO", "commandstats")
	dbSize, _ := h.redis.DbSize()

	infoMap := parseRedisInfo(fmt.Sprintf("%v", infoRaw))
	cmdStats := parseCommandStats(fmt.Sprintf("%v", cmdRaw))

	result := model.CacheInfoResult{
		Info:         infoMap,
		DbSize:       dbSize,
		CommandStats: cmdStats,
	}

	ctx.JSON(common.Success(result))
}

func parseRedisInfo(raw string) map[string]string {
	result := make(map[string]string)
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}
	return result
}

func parseCommandStats(raw string) []model.CommandStat {
	var stats []model.CommandStat
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || !strings.HasPrefix(line, "cmdstat_") {
			continue
		}
		name := strings.TrimPrefix(line[:strings.Index(line, ":")], "cmdstat_")
		value := line[strings.Index(line, ":")+1:]
		calls := ""
		if s := strings.TrimPrefix(value, "calls="); s != "" {
			if idx := strings.Index(s, ","); idx > 0 {
				calls = s[:idx]
			} else {
				calls = s
			}
		}
		stats = append(stats, model.CommandStat{Name: name, Value: calls})
	}
	return stats
}

func (h *MonitorHandler) CacheNames(ctx iris.Context) {
	ctx.JSON(common.Success(cacheNames))
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
	result := model.SysCache{
		CacheName:  cacheName,
		CacheKey:   cacheKey,
		CacheValue: val,
	}
	ctx.JSON(common.Success(result))
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
	cacheKey := ctx.Params().GetString("cacheKey")
	h.redis.Del(cacheKey)
	ctx.JSON(common.Success(nil))
}

func (h *MonitorHandler) ClearCacheAll(ctx iris.Context) {
	keys, _ := h.redis.Keys("*")
	for _, key := range keys {
		h.redis.Del(key)
	}
	ctx.JSON(common.Success(nil))
}
