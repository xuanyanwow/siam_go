package core

import (
    "github.com/astaxie/beego/config"
    "github.com/kataras/iris/v12"
)

var (
    // 配置管理器
    S_Appconf config.Configer
    // 主Iris框架服务
    S_Iris *iris.Application
)