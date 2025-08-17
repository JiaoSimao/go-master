package utils

import (
	"fmt"
	"github.com/go-master/task4/model"
	"gopkg.in/yaml.v3"
	"os"
)

var config *model.Config

// LoadConfig 加载配置文件内容到结构体Config中
func LoadConfig(path string) (*model.Config, error) {
	//读取配置文件内容
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file %s error: %v", path, err)
	}
	//创建一个空的config，yaml.Unmarshal需要
	config = &model.Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal file data %s error: %v", path, err)
	}
	if config.Mysql.DSN == "" {
		return nil, fmt.Errorf("dsn is empty")
	}
	return config, nil
}
