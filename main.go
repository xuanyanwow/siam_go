package main

import (
	"main/config"
	"main/core"
)

func main() {
	// 初始化
	core.InitFrameWork();
	// 加载中间件
	config.Middleware()
	// 加载路由
	config.Router();
	// 开始监听
	core.StartListen();
}

