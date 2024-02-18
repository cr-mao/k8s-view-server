package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/middlewares"
	"github.com/cr-mao/k8s-view-server/infra/app"
	"github.com/cr-mao/k8s-view-server/infra/conf"
	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"
)

// 404处理
func setup404Handler(r *gin.Engine) {
	// 添加 Get 请求路路由
	r.NoRoute(func(c *gin.Context) {
		response.ErrorAbort(c, errcode.ErrCodes.ErrNotFound)
	})
}

//全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),   //自定义请求响应中间件
		middlewares.Recovery(), //panic   错误 拦截处理
		middlewares.Cros(),     //跨域
	)
}

func NewRouter() *gin.Engine {
	if app.IsLocal() && conf.GetBool("app.app_debug", false) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	registerGlobalMiddleWare(router)

	setup404Handler(router)
	//外部api
	RegisterAPIRoutes(router)
	return router
}
