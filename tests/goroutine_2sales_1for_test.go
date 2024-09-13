// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com
package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// 商品列表
var products = map[string]*Product{
	"苹果": {stock: 50},
	"栗子": {stock: 30},
}

// 下单函数
func placeOrder(items map[string]int) string {
	// 逐个检查库存并扣减
	for productID, quantity := range items {
		product := products[productID]
		product.mu.Lock()
		defer func() {
			fmt.Printf("解锁：%v\n", productID)
			product.mu.Unlock()
		}()
		if product.stock-quantity >= 0 {
			product.stock -= quantity
			product.sales += quantity
			fmt.Printf("%v下单成功，剩余库存：%d\n", productID, product.stock)
		} else {
			fmt.Printf("%v下单失败，剩余库存：%d\n", productID, product.stock)
			return "商品库存不足"
		}
	}
	return "已下单成功"
}

func TestGoroutine2Sales1For(t *testing.T) {
	var i = 0
	// 模拟1000个下单请求
	for {
		i++
		fmt.Printf("并发下单请求i：%d\n", i)
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
		go func(items map[string]int) {
			placeOrder(items)
		}(items)
	}
}
