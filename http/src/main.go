package main

import (
	"./conf"
	"fmt"
)

func main() {
	//导入配置文件
	configMap := conf.InitConfig("D:/workspace/golang/myfirstgo/http2/src/http_config.conf")
	//获取配置里host属性的value
	fmt.Println(configMap["host"])
	//查看配置文件里所有键值对
	fmt.Println(configMap)
}