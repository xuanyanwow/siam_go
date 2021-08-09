package cron

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)

// Redis 订阅发布demo
func CronTestRedis(){
    redisConnect,err := redis.DialURL("redis://127.0.0.1:6379")
    if err != nil {
        panic(err)
    }
    defer redisConnect.Close()

    psc := redis.PubSubConn{Conn: redisConnect}
    psc.Subscribe("siam_test")
    for {
        switch v := psc.Receive().(type) {
        case redis.Message:
            // 解析数据 执行任务
            fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
        case redis.Subscription:
            fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
        case error:
            panic(v)
        }
    }
}