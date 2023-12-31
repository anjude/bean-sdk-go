package config

import "time"

type Config struct {
	BaseUrl string
	Timeout time.Duration
}

func GetDefaultConfig() *Config {
	return &Config{
		BaseUrl: "https://golang-fkr4-1783471-1303969980.ap-shanghai.run.tcloudbase.com",
	}
}
