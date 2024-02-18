/**
User: cr-mao
Date: 2024/02/18 14:06
Email: crmao@qq.com
Desc: bootstrap.go
*/
package bootstrap

import (
	"math/rand"
	"time"

	"github.com/cr-mao/k8s-view-server/infra/conf"
)

func Bootstrap(env string) {
	//  配置初始化
	conf.InitConfig(env)
	//随机数种子
	rand.Seed(time.Now().UnixNano())
	//全局设置时区
	var cstZone, _ = time.LoadLocation(conf.GetString("app.timezone"))
	time.Local = cstZone
	// logger 初始化
	SetupLogger()
}
