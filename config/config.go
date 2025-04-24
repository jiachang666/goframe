package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Conf struct {
	Nacos
}

type Nacos struct {
	NameId string
	DataId string
	Group  string
	Host   string
	Port   int
}

var Config Conf

func Init() {
	viper.SetConfigFile("../common/config/dev.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("配置文件读取失败")
		return
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		fmt.Println("配置文件解析失败")
		return
	}
	fmt.Println("viper初始化成功")

}
