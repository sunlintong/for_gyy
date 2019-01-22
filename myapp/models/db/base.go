package db

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
	"sync"
)

// 初始化数据库连接
// func init() {
// 	sql := beego.AppConfig.String("mysqluser") + ":" +
// 		beego.AppConfig.String("mysqlpassword") + "@tcp(" +
// 		beego.AppConfig.String("mysqlhost") + ":" +
// 		beego.AppConfig.String("mysqlport") + ")/" +
// 		beego.AppConfig.String("mysqldatabase") + "?charset=utf8"
	
// 	orm.RegisterDataBase("default", "mysql", sql, 30)
// 	orm.RegisterModel(new(User), new(Log), new(Container), new(Image))
// 	orm.RunSyncdb("default", false, true)
// }

var (
	globalOrm orm.Ormer
	once      sync.Once
)

// 单例模式获取数据库ormer
func GetOrmer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
