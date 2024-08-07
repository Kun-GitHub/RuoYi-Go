// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package ryi18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// 获取Localizer实例
func LoadLocalizer(lang string) *i18n.Localizer {
	var bundle = i18n.NewBundle(language.SimplifiedChinese)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// 使用LoadMessageFile加载YAML文件，并检查错误
	bundle.LoadMessageFile("./locales/en-US.yaml")
	bundle.LoadMessageFile("./locales/zh-CN.yaml")

	if bundle != nil {
		return i18n.NewLocalizer(bundle, lang)
	} else {
		return nil
	}
}
