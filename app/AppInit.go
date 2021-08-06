package app

import (
	"main/app/utils"
	"main/config"
	"main/core"
)

func Init(){
	LoadDefaultEvent()
	// 注册数据库连接池
	utils.GormPoolInit()

}

func LoadDefaultEvent(){
	core.RegisterEvent(core.EventNameBeforeStartListen, config.OnBeforeStartListen)
	core.RegisterEvent(core.EventNameOnShutDown, config.OnShutDown)
}