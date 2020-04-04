package util

import (
	"myList/model"

	"github.com/jinzhu/gorm"
)

func IsEmailExist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email=?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
