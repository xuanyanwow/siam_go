package config

import (
    "github.com/kataras/iris/v12"
)

func OnBeforeRequest(ctx iris.Context){
    //ctx.Application().Logger().Info("test siam")
    //log.Println("事件测试,beforeRequest")
    ctx.Next()
}