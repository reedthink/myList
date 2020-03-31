package middleware

import (
	"auth/dao"
	"auth/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header

		tokenString := c.GetHeader("authorization")
		log.Printf("%v gg.",tokenString)

		log.Println(tokenString)
		//验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足1",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足2",
			})
			c.Abort()
			return
		}
		//通过验证，获取userId
		userId:=claims.UserId
		DB:=dao.GetDB()
		var user model.User
		DB.First(&user,userId)
		//用户不存在
		if user.ID==0{
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足3",
			})
			c.Abort()
			return
		}

		//用户存在,将user信息写入上下文
		c.Set("user",user) //todo what?

		c.Next()
	}
}
