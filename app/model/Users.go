package model

import (
    "time"
)

type UsersModel struct {
    U_Id          uint64			`gorm:"primaryKey"`
    U_Name         string
    Create_Time   time.Time
}

func (UsersModel) TableName() string{
    return "users"
}