package controller

import (
	"gorm.io/gorm"
	"main/app/model"
	"main/app/utils"
)

func Login() model.UsersModel {
	db := utils.GetDi(utils.DiNameDefaultDb).(*gorm.DB)
	userModel := model.UsersModel{}
	db.Where("u_name = ?", "siam").First(&userModel)
	//time.Sleep(time.Second * 15)
	return userModel
}
