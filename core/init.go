package core

import (
	"github.com/astaxie/beego/config"
	"github.com/kataras/iris/v12"
	"log"
)

func InitFrameWork() {
	LoadConf()

	app := iris.Default()
	S_Iris = app;


	app.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			println("server terminated")
		})
	})

}
func StartListen() {
	S_Iris.Listen(":8080");
}


//从配置文件读取配置信息
func LoadConf() {
	appConf, err := config.NewConfig("ini", "./config/main.conf")
	if err != nil {
		log.Println(err)
		return
	}
	S_Appconf = appConf;
}
