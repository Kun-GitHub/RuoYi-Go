// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package main

import (
	"RuoYi-Go/pkg/captcha"
	"testing"
)

func TestGenerateCaptcha(t *testing.T) {
	id, b64s, a, _ := captcha.GenerateCaptcha()
	println(b64s)
	b := captcha.VerifyCaptcha(id, a)
	println(b)
}
