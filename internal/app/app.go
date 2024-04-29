package app

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
)

// Run 初始化并运行应用
func Run(config interface{}, localizer *i18n.Localizer, log *zap.Logger) {
	// 假设的一些应用初始化逻辑...

	// 使用本地化信息输出欢迎信息
	welcomeMessage := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "welcome",
			Other: "Welcome to our application!",
		},
	})
	fmt.Println(welcomeMessage)

	// 应用的其他逻辑...
}
