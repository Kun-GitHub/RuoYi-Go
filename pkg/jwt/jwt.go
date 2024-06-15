package ryjwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const key = "j17GjwcQfeFVDxlSx7RW"

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
