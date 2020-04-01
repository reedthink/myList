package main

import (
	"auth/controller"
	"auth/middleware"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	r.Static("/static", "static") //指定样式文件地址
	r.LoadHTMLGlob("templates/*") //指定静态文件位置
	r.Use(m1)                     //全局注册中间件函数
	//r.Use(middleware.AuthMiddleware())
	r.GET("/list", controller.IndexHandler)
	//API v1
	v1Group := r.Group("v1")
	{
		//todo list
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}

//中间件,统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Printf("m1 out.\n")
}
