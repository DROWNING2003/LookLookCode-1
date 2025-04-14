package initial

import (
	"fmt"
	"log"

	"base/config"

	"github.com/spf13/viper"
)

// InitViper 初始化配置
func InitViper() {
	viper.SetConfigName("config") // 配置文件名称
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath("config") // 配置文件路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 将配置映射到结构体
	if err := viper.Unmarshal(&config.GlobalConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	fmt.Println("Configuration loaded successfully")
}
