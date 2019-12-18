package main

import (
	dbutil2 "dbutil"
	router2 "router"
)

func main() {
	//加载配置文件并初始化连接
	db := dbutil2.InitCon("./config.ini", "mysql")
	//初始化路由
	router2.InitRouter()
	//关闭数据库连接
	defer dbutil2.CloseCon(db)
}
