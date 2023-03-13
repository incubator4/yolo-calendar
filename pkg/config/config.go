package config

import "github.com/spf13/viper"

type Config struct {
	DbUrl                 string
	TencentCloudSecretId  string
	TencentCloudSecretKey string
}

var GlobalConfig Config

func init() {
	viper.SetEnvPrefix("yolo")
	_ = viper.BindEnv("DB_URL")
	_ = viper.BindEnv("TENCENTCLOUD_SECRET_ID")
	_ = viper.BindEnv("TENCENTCLOUD_SECRET_KEY")
	GlobalConfig.DbUrl = viper.GetString("DB_URL")
	GlobalConfig.TencentCloudSecretId = viper.GetString("TENCENTCLOUD_SECRET_ID")
	GlobalConfig.TencentCloudSecretKey = viper.GetString("TENCENTCLOUD_SECRET_KEY")
}
