package config

import (
	"main/app/middleware"
	"main/core"
)

// 注入全局的中间件事件
func Middleware(){
	//core.S_Iris.Use(OnBeforeRequest)
	//core.S_Iris.Use(OnAfterRequest)
	core.S_Iris.Use(middleware.CorsMiddleware)
}
