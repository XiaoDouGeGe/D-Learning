package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// 服务端配置
type AppConfig struct {
	ServerPort  string `json:"server_port" yaml:"server_port"`
	ResourceDir string `json:"resource_dir" yaml:"resource_dir"`
	DBDriver    string `json:"db_driver" yaml:"db_driver"`
	DBSource    string `json:"db_source" yaml:"db_source"`
}

// 初始化服务器配置
func InitConfig() *AppConfig {
	var config *AppConfig
	content, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err.Error())
	}
	return config
}
