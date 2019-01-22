package main

import (
	// 导入routers包是为了执行包中的init函数来初始化路由
	_ "github.com/sunlintong/for_gyy/myapp/routers"

	"github.com/astaxie/beego"
)

// 入口main函数
func main() {
	beego.Run()
}
