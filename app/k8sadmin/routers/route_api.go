package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/controllers"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	loginController := &controllers.AdminLoginController{}
	//登录
	r.POST("/admin/user/login", loginController.Login)
}
