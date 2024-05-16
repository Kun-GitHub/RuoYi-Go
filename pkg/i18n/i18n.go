// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package ryi18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var bundle *i18n.Bundle

var Localizer *i18n.Localizer

func init() {
	bundle = i18n.NewBundle(language.SimplifiedChinese)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// 使用LoadMessageFile加载YAML文件，并检查错误
	bundle.LoadMessageFile("./locales/en-US.yaml")
	bundle.LoadMessageFile("./locales/zh-CN.yaml")
}

// 获取Localizer实例
func GetLocalizer(lang string) (*i18n.Localizer, error) {
	if bundle != nil {
		Localizer = i18n.NewLocalizer(bundle, lang)
		return Localizer, nil
	} else {
		return nil, nil
	}
}
