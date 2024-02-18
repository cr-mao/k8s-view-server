package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/controllers"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	tmpController := &controllers.TmpController{}
	//PodList 测试
	r.POST("/kubernetes/pod_list", tmpController.PodList)
}
