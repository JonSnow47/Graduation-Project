# Note
注意事项: 
- 管理员注册路默认为注释状态,不对外暴露  `/blog/routers/router.go line:10`
- API 接收的文章内容应为 **base64编码字符串**  `/blog/controllers/article.go line:41`
- 