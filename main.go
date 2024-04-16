package main

import (
	"machinesearch/bootstrap"
	"machinesearch/global"
	"machinesearch/router"
)

func main() {

	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()

	//启动定时任务
	go bootstrap.InitializeCron()

	//启动gin服务
	router.RunServer()

}
