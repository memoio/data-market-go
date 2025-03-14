package config

import (
	"encoding/json"
	"os"
)

var Cfg *Config

type Config struct {
	LogLevel string `json:"log_level"`
}

// Init 从指定路径初始化配置（默认 config.json）
func init() {
	// default path
	path := "config.json"

	// open config.json file
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cfg := &Config{
		LogLevel: "info", // 默认值
	}
	if err := json.NewDecoder(file).Decode(cfg); err != nil {
		panic(err)
	}

	Cfg = cfg
}
