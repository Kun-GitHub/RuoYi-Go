// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package main

import (
	ryjwt "RuoYi-Go/pkg/jwt"
	"testing"
)

func TestSing(t *testing.T) {
	s, _ := ryjwt.Sign("id", "1234", 1)
	v, _ := ryjwt.Valid("id", s)
	println(v)
}
