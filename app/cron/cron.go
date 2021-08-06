package cron

import (
	"github.com/robfig/cron/v3"
	"log"
)

/**
 项目Cron事务管理
 */
func Start(){
	log.Println("Starting...")
	c := cron.New(cron.WithSeconds())




	c.AddFunc("* * * * * *", func() {
		log.Println("Run...")
	})




	go c.Start()
	defer c.Stop()
	select {}
}