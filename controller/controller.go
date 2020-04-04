package controller

import (
	"github.com/gin-gonic/gin"
	"myList/dao"
	"myList/model"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
//func Info(c *gin.Context) {
//	user, _ := c.Get("user")
//	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{
//		"user": dto.ToUserDto(user.(model.User)), //user是一个接口
//	}})
//}
func CreateATodo(c *gin.Context) {
	user, _ := c.Get("user")
	email := user.(model.User).Email
	//1.从请求中把数据拿出来
	var todo model.Todo
	c.BindJSON(&todo)
	//var requestUser model.User
	//c.BindJSON(&requestUser)
	todo.Email = email
	//2.存入数据库并且返回响应
	err := dao.DB.Create(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todo,
		})
	}
}

//查询todo数据表
func GetTodoList(c *gin.Context) {
	user, _ := c.Get("user")
	email := user.(model.User).Email
	todoList, err := GetTodoList_A(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList) //返回json
	}

}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的ID",
		})
		return
	}
	todo, err := GetATodoA(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
	if err = UpdateATodoA(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
		return
	}
	err := DeleteATodoA(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}

/*todo 下方为具体实现的函数，之后改写*/
//GetTodoList 查询单个用户所有todo
func GetTodoList_A(email string) (todoList []*model.Todo, err error) {
	err = dao.DB.Where("email=?", email).Find(&todoList).Error
	if err != nil {
		return todoList, err
	}
	return
}

//GetATodo 查询单个用户指定的todo
func GetATodoA(id string) (*model.Todo, error) {
	todo := new(model.Todo)
	err := dao.DB.Where("id=?", id).First(todo).Error
	if err != nil {
		return nil, err
	}
	return todo, err

}
func UpdateATodoA(todo *model.Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}
func DeleteATodoA(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&model.Todo{}).Error
	return
}
