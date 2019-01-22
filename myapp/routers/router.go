package routers

import (
	"github.com/astaxie/beego"
	"github.com/sunlintong/for_gyy/myapp/controllers"
)

// 路由过滤器
// var FilterUser = func(ctx *context.Context) {
// 	user, ok := ctx.Input.Session("user").(types.User)
// 	url := ctx.Request.RequestURI
// 	log.Printf("url: %s, session: %+v, pass?: %v\n", url, user, ok)
// 	// 没有session说明用户还未登录，这时如果请求其他页面，则跳转至登录页面
// 	if !ok && url != "/login" {
// 		if url == "/register" {
// 			return
// 		}
// 		ctx.Redirect(302, "login")
// 	}
// }

func init() {
	// 对用户的请求进行过滤，
	// beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

	// 登录注册 路由
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
}
