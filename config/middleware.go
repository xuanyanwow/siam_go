package config

import (
	"github.com/kataras/iris/v12"
	"main/app/event"
	"main/core"
)


// 注入中间件事件
func Middleware(){
	core.S_Iris.Use(myMiddleware);
	core.S_Iris.Use(event.OnAfterRequest)
}


// 中间件1

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("运行之前 %s", ctx.Path())
	ctx.Next()
}