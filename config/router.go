package config

import (
	"github.com/kataras/iris/v12"
	"main/app/controller"
	"main/core"
)

/*
  注册路由
 */
func Router(){
	core.S_Iris.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})
	core.S_Iris.Handle("GET", "/user/login", func(ctx iris.Context) {
		ctx.JSON(controller.Login());
	})
}