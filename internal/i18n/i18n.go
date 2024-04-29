package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var bundle *i18n.Bundle

func InitializeI18n() error {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// 使用LoadMessageFile加载YAML文件，并检查错误
	_, err := bundle.LoadMessageFile("./locales/en-US.yaml")
	if err != nil {
		return err
	}
	_, err = bundle.LoadMessageFile("./locales/zh-CN.yaml")
	if err != nil {
		return err
	}
	return nil
}

// 获取Localizer实例
func GetLocalizer(lang string) (*i18n.Localizer, error) {
	return i18n.NewLocalizer(bundle, lang), nil
}
