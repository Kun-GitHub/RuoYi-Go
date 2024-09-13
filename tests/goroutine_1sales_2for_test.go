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

// 某个商城下单单个商品的情况，模拟1000个并发

func TestGoroutine1Sales2For(t *testing.T) {
	// 初始化商品
	product := &Product{stock: 50}

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 创建一个等待组，用于等待所有 goroutines 完成
	var wg sync.WaitGroup

	// 模拟1000个下单请求，分为两个 goroutines
	quantity := 1 // 假设每次下单只买一个商品

	// 第一个 goroutine 处理 500 个请求
	wg.Add(1)
	go func(quantity int) {
		defer wg.Done()
		for i := 0; i < 500; i++ {
			fmt.Printf("并发下单i：%d\n", i)
			wg.Add(1)
			go func(quantity int) {
				defer wg.Done()
				product.PlaceOrder(quantity)
			}(quantity)
		}
	}(quantity)

	// 第二个 goroutine 处理另外 500 个请求
	wg.Add(1)
	go func(quantity int) {
		defer wg.Done()
		for j := 0; j < 500; j++ {
			fmt.Printf("并发下单j：%d\n", j)
			wg.Add(1)
			go func(quantity int) {
				defer wg.Done()
				product.PlaceOrder(quantity)
			}(quantity)
		}
	}(quantity)

	// 等待所有 goroutines 完成
	wg.Wait()
}
