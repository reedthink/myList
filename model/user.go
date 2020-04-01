package model

import "github.com/jinzhu/gorm"

//User 用户信息模型
type User struct {
	gorm.Model //基本模型
	Name     string `gorm:"type:varchar(20);not null"` //用户名长度20,不为空
	Email    string `gorm:"type:varchar(100);unique_index"`//邮箱长度100,为该列设置唯一索引
	Password string `gorm:"size:255;not null"`
}
//表名（Table Name）
//表名默认就是结构体名称的复数