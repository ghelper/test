package config

import (
	"Project/common/log"
	"Project/common/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func DefaultViper() {
	v := viper.New()
	v.SetConfigName("default")
	v.SetConfigType("toml")

	v.AddConfigPath(utils.GetConfigPath())
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Read Config Error: %s", err.Error())
	}

	// 读取所有配置
	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}
}

func SpecifyViper(configName, configType, configPath string) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read Config Error: %s", err.Error())
	}

	// 配置动态加载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("config file is changed:", e.Name)
	})
}

func DefaultConfig() error {
	v := viper.New()
	v.SetConfigName("default")
	v.SetConfigType("toml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	// 读取所有配置
	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}
	return nil
}
