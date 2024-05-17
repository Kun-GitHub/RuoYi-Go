// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package response

type CaptchaImage struct {
	Code           int    `json:"code"`
	Message        string `json:"msg"`
	Uuid           string `json:"uuid"`
	CaptchaEnabled bool   `json:"captchaEnabled,omitempty"`
	Img            string `json:"img"`
}
