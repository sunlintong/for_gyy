package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

type BaseController struct {
	beego.Controller
}

// 密码加密函数，将密码加盐再MD5
func (bc *BaseController) Encode(password string) string {
	salt := "sunlintonghenshuai"
	h := md5.New()
	h.Write([]byte(password + salt))
	return hex.EncodeToString(h.Sum(nil))
}

// 将unix时间戳格式化为字符串
func (bc *BaseController) GetTimeString(now int64) string {
	return time.Unix(now, 0).Format("2006-01-02 15:04:05")
}

// 错误处理
func (bc *BaseController) CheckErr(err error) {
	if err != nil {
		glog.V(2).Infoln(err)
	}
}

// ReturnJSON return json
func (bc *BaseController) ReturnJSON(resp interface{}) {
	bc.Data["json"] = resp
	bc.ServeJSON()
	bc.Finish()
}

// Success 200
func (bc *BaseController) Success(resp interface{}) {
	bc.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	bc.ReturnJSON(resp)
}

// BadRequest 400 参数验证失败
func (bc *BaseController) BadRequest(resp interface{}) {
	bc.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	bc.ReturnJSON(resp)
}

// ServiceError 500
func (bc *BaseController) ServiceError(resp interface{}) {
	glog.Errorln(resp)
	bc.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	bc.Data["json"] = resp
	bc.ServeJSON()
	bc.Finish()
}
