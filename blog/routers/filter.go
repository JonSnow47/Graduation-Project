package routers

import (
	"github.com/astaxie/beego/context"

	"github.com/JonSnow47/Graduation-Project/blog/consts"
)

var MapFilter map[string]interface{}

func init() {
	MapFilter = make(map[string]interface{})

	MapFilter["/admin/login"] = struct{}{}

	MapFilter["/blog/article/get"] = struct{}{}
	MapFilter["/blog/article/approved"] = struct{}{}

	MapFilter["/blog/tag/get"] = struct{}{}
	MapFilter["/blog/tag/all"] = struct{}{}

}

func LoginFilter(ctx *context.Context) {
	if _, ok := MapFilter[ctx.Request.RequestURI]; !ok {
		id := ctx.Input.CruSession.Get(consts.SessionId)
		if id == nil {
			//ctx.Output.JSON(map[string]interface{}{consts.Status: consts.ErrLoginRequired}, false, false)
			ctx.Redirect(302, "/admin/login")
		}
	}
}
