// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package main

import (
	"RuoYi-Go/di"
	"RuoYi-Go/internal/shutdown"
	"RuoYi-Go/pkg/config"
	"os"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		os.Exit(2)
	}

	// 创建依赖注入容器
	container, err := di.NewContainer(cfg)
	if err != nil {
		os.Exit(2)
	}
	defer container.Close()

	container.InitJob()
	err = container.StartServer()
	if err != nil {
		os.Exit(2)
	}

	// 系统关闭
	shutdown.NewHook().Close(
		func() {
			container.Close()
			os.Exit(0)
		},
	)
}
