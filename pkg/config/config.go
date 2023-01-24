package config

import "github.com/spf13/viper"

type Config struct {
	DbUrl string
}

var GlobalConfig Config

func init() {
	viper.SetEnvPrefix("yolo")
	_ = viper.BindEnv("DB_URL")
	GlobalConfig.DbUrl = viper.GetString("DB_URL")
}
