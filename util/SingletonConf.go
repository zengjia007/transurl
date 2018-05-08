package util

import (
	"transurl/confutil"
	"sync"
)

type singletonConf struct {
	c *confutil.Config
}

var instance *singletonConf

var mutex sync.Mutex

// 使用单例模式获取confutil.Config对象
// 提供对外获取单例配置解析对象
func GetInstanceConf() *confutil.Config {
	mutex.Lock()
	defer mutex.Unlock()
	if instance == nil {
		instance = &singletonConf{}
		instance.c  = confutil.InitConfig("resources/default.properties")
	}
	return instance.c
}
