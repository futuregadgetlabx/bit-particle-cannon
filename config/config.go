package config

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Application struct {
	Lark    *Lark    `yaml:"lark" mapstructure:"lark"`
	TianApi *TianApi `yaml:"tianapi" mapstructure:"tianapi"`
	Cron    string   `yaml:"cron" mapstructure:"cron"`
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
	v.AddConfigPath("/app/config")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		}
		log.Fatal(err)
	}

	App = &Application{}
	err := v.Unmarshal(App)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		App = &Application{}
		err = v.Unmarshal(App)
		if err != nil {
			logrus.WithError(err).Fatal("config update error")
		}
	})
	if os.Getenv("ENV") == "dev" {
		v.Debug()
		logrus.Debugf("%+v", *App)
	}
	if err != nil {
		panic(err)
	}
}
