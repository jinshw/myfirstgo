package main

import (
	"./common"
	"./conf"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//导入配置文件
	configMap := conf.InitConfig("http_config.conf")
	//获取配置里host属性的value
	fmt.Println(configMap["host"])
	//查看配置文件里所有键值对
	fmt.Println(configMap)

	port := ":" + configMap["port"]

	router := common.NewRouter(configMap)
	fmt.Println("http service start success!port" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
