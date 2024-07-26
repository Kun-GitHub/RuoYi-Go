// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"testing"
)

func TestServer(t *testing.T) {
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

}
