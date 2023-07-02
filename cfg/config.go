package cfg

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config *viper.Viper

func init() {
	config = viper.New()
	//优先读取根目录配置文件
	config.SetConfigName("config")  // 配置文件名，不需要后缀名
	config.SetConfigType("yml")     // 配置文件格式
	config.AddConfigPath(".")       // 查找配置文件的路径
	config.AddConfigPath("config/") // 查找配置文件的路径
	err := config.ReadInConfig()    // 查找并读取配置文件
	if err != nil {                 // 处理错误
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file not found"))
		}
	}
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	config.WatchConfig()
}

func Get(key string) interface{} {
	return config.Get(key)
}

func GetStr(key string) string {
	return config.GetString(key)
}
