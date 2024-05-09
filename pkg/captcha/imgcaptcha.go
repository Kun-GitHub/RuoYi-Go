// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package main

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"log"
	"net/http"
)

func main() {
	// 定义验证码配置
	var store = base64Captcha.DefaultMemStore
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	c := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	a, id, b64s, err := c.Generate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated Code ID:", a)
	fmt.Println("Generated Code ID:", id)
	fmt.Println("Generated Base64 String:", b64s)

	// 如果要在Web应用中使用，可以这样发送验证码到前端
	http.HandleFunc("/captcha", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(b64s)) // 发送Base64编码的图片到前端
	})

	// 启动HTTP服务器以便测试（实际应用中可能不需要这一步）
	log.Println("Server started on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
