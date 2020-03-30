package controller

import (
	"auth/dao"
	"auth/dto"
	"auth/middleware"
	"auth/response"
	//"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"auth/model"
	"auth/util"

	"github.com/gin-gonic/gin"
)

//Register  注册函数
func Register(c *gin.Context) {
	db := dao.GetDB()
	//获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	name := requestUser.Name
	email := requestUser.Email
	password := requestUser.Password

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	if len(name) == 0 {
		name = "HanPi"
	}

	log.Println(name, email, password)
	//判断邮箱是否存在
	if util.IsEmailExist(db, email) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "邮箱已被注册,GG")
		return
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			response.Response(c, http.StatusInternalServerError, 500, nil, "加密失败")

			return
		}
		newUser := model.User{
			Name:     name,
			Email:    email,
			Password: string(hashedPassword),
		}
		db.Create(&newUser)
		token, _:= middleware.ReleaseToken(newUser)
		response.Response(c, http.StatusOK, 200, gin.H{"token": token}, "注册成功")
	}
}

func Login(c *gin.Context) {
	DB := dao.GetDB()
	//获取参数
	var requestUser = model.User{}
	c.Bind(&requestUser)
	email := requestUser.Email
	password := requestUser.Password
	//数据验证
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	//判断邮箱是否存在
	var user model.User
	DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	// 若正确，发放token
	token, err := middleware.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		//记录日志
		log.Printf("tolen generate err: %v", err)
		return
	}
	//返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{
		"user": dto.ToUserDto(user.(model.User)), //user是一个接口
	}})
}
