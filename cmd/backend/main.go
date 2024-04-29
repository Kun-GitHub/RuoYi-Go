package main

import (
	"time-machine/internal/app"
	"time-machine/internal/config"
	"time-machine/internal/i18n"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	if err := i18n.InitializeI18n(); err != nil {
		panic(err)
	}

	localizer, err := i18n.GetLocalizer("zh-CN")
	if err != nil {
		panic(err)
	}

	app.Run(cfg, localizer)
}
