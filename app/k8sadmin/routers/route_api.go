package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/controllers"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	group := r.Group("/k8s")
	namespaceController := &controllers.NamespaceController{}
	podController := &controllers.PodController{}

	// 命名空间列表
	group.GET("/namespace_list", namespaceController.GetNamespaceList)
	// pod 创建更新
	group.POST("/pod_create_update", podController.CreateOrUpdate)
	// 删除pod
	group.GET("/pod_delete/:namespace/:name", podController.DeletePod)
}
