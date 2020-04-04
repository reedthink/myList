云端待办清单的后端。  
go语⾔编写，实现了注册，登录，记录待办事项的功能。每个账户有独⾃的待办事项。  
RESTful的API，采用token进行用户认证  
版本要求: go1.13以上  
数据库 ： mysql v8  
配置文件存放位置： `./config/application.yml`

## 查看API文档: http://114.116.239.137:8080/swagger/index.html

## 如何启动: `go build && ./myList`

- web框架：gin
- ORM：gorm
- json web token : jwt-go
- 配置文件加载：viper
- api文档生成: swaggo
- 密码采用crypto库bcrypt包加密


---
备忘：
1. 函数长度控制在一屏内
2. 关于gomod导入当前文件夹下的包： 导入当前主包的模块名和子文件夹的包名即可