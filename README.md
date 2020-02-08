#### 运行服务
```
go build -o output/bin/gin-learn-todo  app/main.go

./output/bin/gin-learn-todo

```

#### 参考资料

```
Gin框架 学习
https://www.kancloud.cn/shuangdeyu/gin_book/949445

Gin学习文档
https://learnku.com/docs/gin-gonic/2019/examples-jsonp/6162

```
#### Swagger
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles

...
在项目的根目录：windows下执行swag.exe init ，linux下执行swag init

参考地址: https://blog.csdn.net/hjxzb/article/details/84899100

go build -tags=doc 

```