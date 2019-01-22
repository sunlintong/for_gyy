package controllers

import (
	"encoding/json"
	"log"
)

// 实现BaseController，能方便的调其中函数，简化响应操作
type RegisterController struct {
	BaseController
}

type RegisterRequest struct {
	User_Name     string `json:"user_name"`
	User_Password string `json:"user_password"`
}

// GET方法展示注册页面
func (rc *RegisterController) Get() {
	rc.TplName = "register.html"
}

// POST方法处理注册请求
func (rc *RegisterController) Post() {
	var req RegisterRequest
	// 将请求的body解析到req这个对象，后面对这个对象进行操作就行啦
	json.Unmarshal(rc.Ctx.Input.RequestBody, &req)
	log.Println("收到注册请求，进行逻辑处理", req)

	// 这里就可以进行处理balabala的了

	// 最后返回数据即可，如果错了就用BadRequest或者ServiceError方法
	rc.Success("注册成功！")
}
