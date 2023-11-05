package config

import (
	"github.com/spf13/viper"
)

type Application struct {
	Lark    *Lark    `yaml:"lark" mapstructure:"lark"`
	TianApi *TianApi `yaml:"tianapi" mapstructure:"tianapi"`
}

type Lark struct {
	Bot    *Bot   `yaml:"bot" mapstructure:"bot"`
	ChatID string `yaml:"chat_id" mapstructure:"chat_id"`
}

type Bot struct {
	AppId     string `yaml:"app_id" mapstructure:"app_id"`
	AppSecret string `yaml:"app_secret" mapstructure:"app_secret"`
}

type TianApi struct {
	Key string `yaml:"key" mapstructure:"key"`
}

var App *Application

func Load() {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	// 读取配置文件
	err := v.ReadInConfig()
	App = &Application{}
	err = v.Unmarshal(App)
	if err != nil {
		panic(err)
	}
}
