// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package model

type CaptchaImage struct {
	Code           int    `json:"code"`
	Message        string `json:"msg"`
	Uuid           string `json:"uuid"`
	CaptchaEnabled bool   `json:"captchaEnabled,omitempty"`
	Img            string `json:"img"`
}
