# Graduation-Project

  Graduation Project: Design a blog background system based on MongoDB
  
  基于 MongoDB 的个人博客后台系统设计

## 后台系统编译需求

- Go 环境配置

  Go 官网: [Go](http://golang.org)
  
  [下载](https://golang.org/dl/)适合自己机器的 Go 版本
  
  环境配置教程: [环境配置](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.0.md)
  
### 步骤

- 下载源代码

    ```bash
    go get github.com/JonSnow47/Graduation-Project
    ```
    
- 编译

    ```bash
    go build
    ./blog
    ```
    
    如果命令窗显示 `2018/05/19 20:14:09.299 [I] [asm_amd64.s:2361] http server Running on http://:8080
` 就表示后台系统启动成功了，你可以通过访问 [localhost:8080](localhost:8080) 来访问后台 API

### API 提供的方法

    ```
	/admin/new      Method: Post
	/admin/login    Method: Post
	/admin/logout   Method: Get

	/blog/article/new           Method: Post
	/blog/article/delete        Method: Post
	/blog/article/update        Method: Post
	/blog/article/modifystate   Method: Post
	/blog/article/get           Method: Get
	/blog/article/all           Method: Post
	/blog/article/approved      Method: Get
	/blog/article/created       Method: Post

	/blog/tag/new       Method: Post
	/blog/tag/delete    Method: Post
	/blog/tag/enable    Method: Post
	/blog/tag/get       Method: Get
	/blog/tag/all       Method: Post
	```