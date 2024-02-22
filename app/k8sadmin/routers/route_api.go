package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/controllers"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	group := r.Group("/k8s")
	// 命名空间
	namespaceController := &controllers.NamespaceController{}
	// 命名空间列表
	group.GET("/namespace_list", namespaceController.GetNamespaceList)
}
