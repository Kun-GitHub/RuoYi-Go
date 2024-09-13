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

// 商品列表
var products = map[string]*Product{
	"苹果": {stock: 50},
	"栗子": {stock: 30},
}

// 下单函数
func placeOrder(items map[string]int) {
	// 逐个检查库存并扣减
	for productID, quantity := range items {
		product := products[productID]
		product.mu.Lock()
		defer product.mu.Unlock()

		if product.stock-quantity >= 0 {
			product.stock -= quantity
			product.sales += quantity

			fmt.Printf("%v下单成功，剩余库存：%d\n", productID, product.stock)
		} else {
			fmt.Printf("%v下单失败，剩余库存：%d\n", productID, product.stock)
		}
	}
}

func TestGoroutine2Sales1For(t *testing.T) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 创建一个等待组，用于等待所有 goroutines 完成
	var wg sync.WaitGroup

	// 模拟1000个下单请求
	for i := 0; i < 1000; i++ {
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

	// 等待所有 goroutines 完成
	wg.Wait()
}
