package main

import (
    "flag"
    "log"
    "main/app"
    "main/app/cron"
    "main/config"
    "main/core"
)

func main() {
    var runMode string
    flag.StringVar(&runMode, "mode", "server", "运行模式,server|cron")
    flag.Parse()
    // 用第一个命令 来运行command
    if runMode=="server" {
        // 初始化
        core.InitFrameWork()
        // 加载中间件
        config.Middleware()
        // 用户定义初始化事件
        app.Init()
        // 加载路由
        config.Router()
        // 设置websocket
        config.WebSocket()
        // 开启用户cron运行
        go cron.Start()
        // 开始监听
        core.StartListen()
    }else if runMode =="cron" {
        // 单独开启cron运行
        cron.Start()
    }else{
        log.Println("mode参数错误")
    }

}

