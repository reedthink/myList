云端待办清单的后端。前端实在是一窍不通，弄不出来，2333    
go语⾔编写，实现了注册，登录，记录待办事项的功能。每个账户有独⾃的待办事项。  
MVC结构,RESTful的API  
采用token进行用户验证

- web框架：gin
- 调用数据库：gorm
- json web token : jwt-go
- 配置文件加载：viper
- 密码采用crypto库bcrypt包加密


开发原则：
1. 函数长度控制在一屏内

关于gomod导入当前文件夹下的包： 导入当前主包的模块名和子文件夹的包名即可