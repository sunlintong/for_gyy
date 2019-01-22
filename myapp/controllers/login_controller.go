package controllers

import (
	"encoding/json"
	"log"
)

type LoginController struct {
	BaseController
}

type LoginRequest struct {
	User_Name     string `json:"user_name"`
	User_Password string `json:"user_password"`
}

// 登录页面
func (lc *LoginController) Get() {
	lc.TplName = "login.html"
}

// 处理登录操作
func (lc *LoginController) Post() {
	var req LoginRequest
	json.Unmarshal(lc.Ctx.Input.RequestBody, &req)
	log.Println("收到登录请求，进行逻辑处理", req)

	lc.Success("登录成功！")
}
