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

// 商品库存结构体
type Product struct {
	mu    sync.Mutex // 互斥锁
	stock int        // 库存
	sales int        // 销售量

	temp int
}

// 下单函数
func (p *Product) Place1SalesOrder(n int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.temp++
	if p.stock-n >= 0 {
		p.stock -= n
		p.sales += n
		fmt.Printf("已下单成功，剩余库存：%d，下单次数：%d\n", p.stock, p.temp)
		return
	}
	fmt.Printf("下单失败，剩余库存：%d，下单次数：%d\n", p.stock, p.temp)
}

func TestGoroutine1Sales1For(t *testing.T) {
	// 初始化商品
	product := &Product{stock: 50}

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 创建一个等待组，用于等待所有 goroutines 完成
	var wg sync.WaitGroup

	// 模拟1000个下单请求
	for i := 0; i < 1000; i++ {
		quantity := 1 // 假设每次下单只买一个商品
		wg.Add(1)
		go func(quantity int) {
			defer wg.Done()
			product.Place1SalesOrder(quantity)
		}(quantity)
	}

	// 等待所有 goroutines 完成
	wg.Wait()
}
