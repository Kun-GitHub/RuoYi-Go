package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var bundle *i18n.Bundle

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
		return i18n.NewLocalizer(bundle, lang), nil
	} else {
		return nil, nil
	}
}
