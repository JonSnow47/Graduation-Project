package routers

import (
	"github.com/astaxie/beego"

	"github.com/JonSnow47/Graduation-Project/blog/controllers"
)

func init() {
	beego.Router("/admin/new", &controllers.AdminController{}, "Post:New")
	beego.Router("/admin/login", &controllers.AdminController{}, "Post:Login")
	beego.Router("/admin/logout", &controllers.AdminController{}, "Post:Logout")

	beego.Router("/blog/article/new", &controllers.ArticleController{}, "Post:New")
	beego.Router("/blog/article/get", &controllers.ArticleController{}, "Post:Get")
	beego.Router("/blog/article/all", &controllers.ArticleController{}, "Post:All")
	beego.Router("/blog/article/update", &controllers.ArticleController{}, "Post:Update")
	beego.Router("/blog/article/approved", &controllers.ArticleController{}, "Post:Approved")
	beego.Router("/blog/article/delete", &controllers.ArticleController{}, "Post:Delete")

	beego.Router("/blog/tag/new", &controllers.TagController{}, "Post:New")
	beego.Router("/blog/tag/delete", &controllers.TagController{}, "Post:Delete")
	beego.Router("/blog/tag/enable", &controllers.TagController{}, "Post:Enable")
	beego.Router("/blog/tag/get", &controllers.TagController{}, "Post:Get")
	beego.Router("/blog/tag/all", &controllers.TagController{}, "Post:All")
}
