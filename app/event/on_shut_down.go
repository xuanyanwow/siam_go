package event

import "log"

func OnShutDown(e interface{}){
	log.Println("事件测试,服务退出")
}