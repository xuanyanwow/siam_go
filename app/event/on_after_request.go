package event

import (
	"github.com/kataras/iris/v12"
	"log"
)

func OnAfterRequest(ctx iris.Context){
	ctx.Record()
	ctx.Next()
	responseBody := string(ctx.Recorder().Body())
	log.Println("事件测试,afterRequest"+responseBody)
}