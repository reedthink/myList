package models

import (
	"auth/dao"
)

//Todo model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}
/*
Todo 增删改查
这个model的增删查改操作都放这里
*/

//CreateATodo 创建todo
func CreateATodo(todo *Todo)(err error){
	err=dao.DB.Create(&todo).Error;
	return
}
//GetTodoList 查询单个用户所有todo
func GetTodoList()(todoList []*Todo,err error){
	err=dao.DB.Find(&todoList).Error
	if err!=nil{
		return todoList,err
	}
	return

}
//GetATodo 查询单个用户指定的todo
func GetATodo(id string)(*Todo,error){
	todo := new(Todo)
	err := dao.DB.Where("id=?",id).First(todo).Error
	if err!=nil{
		return nil,err
	}
	return todo,err

}
func UpdateATodo(todo *Todo)(err error){
	err=dao.DB.Save(todo).Error
	return
}
func DeleteATodo(id string)(err error){
	err=dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}