package bootstrap

import (
	"github.com/cr-mao/k8s-view-server/infra/conf"
	"github.com/cr-mao/k8s-view-server/infra/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		conf.GetString("log.filename"),
		conf.GetInt("log.max_size"),
		conf.GetInt("log.max_backup"),
		conf.GetInt("log.max_age"),
		conf.GetBool("log.compress"),
		conf.GetString("log.log_type"),
		conf.GetString("log.log_level"),
		conf.GetBool("app.app_debug"),
	)
}
