package config

import "log"

func OnBeforeStartListen(e interface{}) {
	log.Println("事件测试,开始监听前");
}