// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package ryjwt

// JWT工具包
// 提供JWT令牌的生成和验证功能

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// JWT签名密钥
const key = "j17GjwcQfeFVDxlSx7RW"

// Sign 生成JWT令牌
// 根据指定的键值对生成带有过期时间的JWT令牌
// 参数:
//   - k: 键名
//   - v: 键值
//   - exp: 过期时间（小时），默认72小时
// 返回值:
//   - string: 生成的JWT令牌
//   - error: 错误信息
func Sign(k, v string, exp int64) (string, error) {
	if exp == 0 {
		exp = 72
	}

	// 定义签名密钥
	signingKey := []byte(key)
	// 创建一个 claims
	claims := jwt.MapClaims{
		k:     v,
		"exp": time.Now().Add(time.Duration(exp) * time.Hour).Unix(), // 设置过期时间
	}

	// 创建一个令牌对象，头部默认类型为 HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥签名令牌
	tokenStr, err := token.SignedString(signingKey)

	//ryredis.Redis.Set(tokenStr, true, time.Duration(exp)*time.Hour)
	return tokenStr, err
}

// Valid 验证JWT令牌
// 验证JWT令牌的有效性并提取指定键的值
// 参数:
//   - k: 要提取的键名
//   - tokenStr: JWT令牌字符串
// 返回值:
//   - string: 提取的键值
//   - error: 错误信息（令牌无效或过期）
func Valid(k, tokenStr string) (string, error) {
	// 定义签名密钥，需要与生成Token时使用的密钥一致
	signingKey := []byte(key)

	// 解析token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//v, err := ryredis.Redis.Get(tokenStr)
		//if err != nil || v == "" {
		//	return "", fmt.Errorf("token失效")
		//}

		return claims[k].(string), nil
	}
	return "", fmt.Errorf("token失效")
}
