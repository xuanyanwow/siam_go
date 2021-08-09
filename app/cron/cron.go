package cron

import (
	"github.com/robfig/cron/v3"
)

/**
 项目Cron事务管理
 */
func Start(){
	c := cron.New(cron.WithSeconds())

	// 定时任务
	//c.AddFunc("* * * * * *", func() {
	//	log.Println("Run...")
	//})

	// 注册长链接监听（类似自定义进程消费者）
	//go CronTestRedis()

	go c.Start()
	defer c.Stop()
	select {}
}