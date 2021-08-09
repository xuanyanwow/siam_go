package utils

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "main/core"
    "time"
)


func GormPoolInit() {
    host := core.S_Appconf.String("mysql::host");
    port := core.S_Appconf.String("mysql::port");
    username := core.S_Appconf.String("mysql::username");
    password := core.S_Appconf.String("mysql::password");
    database := core.S_Appconf.String("mysql::database");
    charset := core.S_Appconf.String("mysql::charset");

    dsn := username+":"+ password +"@tcp("+ host +":"+port+")/"+database+"?charset="+charset+"&parseTime=True&loc=Local"
    db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    sqlDB, _ := db.DB()
    //设置空闲连接池中连接的最大数量
    sqlDB.SetMaxIdleConns(10)
    //设置打开数据库连接的最大数量。
    sqlDB.SetMaxOpenConns(100)
    //设置了连接可复用的最大时间。
    sqlDB.SetConnMaxLifetime(time.Hour)
    RegisterDi(DiNameDefaultDb, db)
}