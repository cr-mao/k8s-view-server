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
	nodeController := &controllers.NodeController{}

	// 命名空间列表
	group.GET("/namespace_list", namespaceController.GetNamespaceList)
	// pod 创建更新
	group.POST("/pod_create_update", podController.CreateOrUpdate)
	// 删除pod
	group.GET("/pod_delete", podController.DeletePod)
	// pod detail

	////node 调度，
	// node详情 /node/detail?node_name=xxx

	group.GET("/node/detail", nodeController.GetNodeDetail)
	// node 列表
	group.GET("/node/node_list", nodeController.GetNodeList)
	// node 打标签,替换的方式
	group.POST("/node/label", nodeController.UpdateNodeLabel)
	//  node污点设置
	// kubectl taint nodes k8snode1 app=app:NoSchedule
	// kubectl taint nodes k8snode1 app=app:NoSchedule-
	group.POST("/node/taint", nodeController.UpdateNodeTaint)
}
