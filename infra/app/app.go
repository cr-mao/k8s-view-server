// Package app 应用信息
package app

import (
	"time"

	"github.com/cr-mao/k8s-view-server/infra/conf"
)

func IsLocal() bool {
	return conf.Get("app.app_env") == "local"
}

func IsProduction() bool {
	return conf.Get("app.app_env") == "production"
}

func IsTesting() bool {
	return conf.Get("app.app_env") == "testing"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	location, _ := time.LoadLocation(conf.GetString("app.timezone"))
	return time.Now().In(location)
}
