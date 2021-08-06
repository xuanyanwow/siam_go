package app

import (
	"main/app/event"
	"main/app/utils"
	"main/core"
)

func Init(){
	LoadDefaultEvent()
	// 注册数据库连接池
	utils.GormPoolInit()

}

func LoadDefaultEvent(){
	core.RegisterEvent(core.EventNameBeforeStartListen, event.OnBeforeStartListen)
	core.RegisterEvent(core.EventNameOnShutDown, event.OnShutDown)
}