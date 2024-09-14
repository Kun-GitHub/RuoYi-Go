// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestGoroutine2Sales2For(t *testing.T) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 创建一个等待组，用于等待所有 goroutines 完成
	var wg sync.WaitGroup

	// 第一个 goroutine 处理 500 个请求
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 500; i++ {
			fmt.Printf("并发下单i：%d\n", i)
			items := make(map[string]int)
			// 随机选择一个商品或两个商品一起下单
			if rand.Intn(2) == 0 {
				// 单个商品
				if rand.Intn(2) == 0 {
					items["苹果"] = 1
				} else {
					items["栗子"] = 1
				}
			} else {
				// 两个商品
				items["苹果"] = 1
				items["栗子"] = 1
			}

			wg.Add(1)
			go func(items map[string]int) {
				defer wg.Done()
				placeOrder(items)
			}(items)
		}
	}()

	// 第一个 goroutine 处理 500 个请求
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 500; j++ {
			fmt.Printf("并发下单j：%d\n", j)
			items := make(map[string]int)
			// 随机选择一个商品或两个商品一起下单
			if rand.Intn(2) == 0 {
				// 单个商品
				if rand.Intn(2) == 0 {
					items["苹果"] = 1
				} else {
					items["栗子"] = 1
				}
			} else {
				// 两个商品
				items["苹果"] = 1
				items["栗子"] = 1
			}

			wg.Add(1)
			go func(items map[string]int) {
				defer wg.Done()
				placeOrder(items)
			}(items)
		}
	}()
	// 等待所有 goroutines 完成
	wg.Wait()
}
