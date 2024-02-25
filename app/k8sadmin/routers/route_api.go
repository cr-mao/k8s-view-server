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
	configMapController := &controllers.ConfigMapController{}
	secretController := &controllers.SecretController{}

	// 命名空间列表
	group.GET("/namespace/list", namespaceController.GetNamespaceList)
	// pod 创建更新
	group.POST("/pod/create_update", podController.CreateOrUpdate)
	// 删除pod
	group.GET("/pod/delete", podController.DeletePod)
	// pod detail

	//******************node调度************************//
	// node详情 /node/detail?node_name=xxx
	group.GET("/node/detail", nodeController.GetNodeDetail)
	// node 列表
	group.GET("/node/list", nodeController.GetNodeList)
	// node 打标签,替换的方式
	group.POST("/node/label", nodeController.UpdateNodeLabel)
	//  node污点设置
	// kubectl taint nodes k8snode1 app=app:NoSchedule
	// kubectl taint nodes k8snode1 app=app:NoSchedule-
	group.POST("/node/taint", nodeController.UpdateNodeTaint)

	//******************ConfigMap************************//
	// 创建或更新configmap
	group.POST("/configmap/create_update", configMapController.CreateOrUpdateConfigMap)
	// configmap详情 ， /configmap/detail?namespace=dev&name=testcm
	group.GET("/configmap/detail", configMapController.GetConfigMapDetail)
	// configmap列表  /configmap/list?namespace=dev
	group.GET("/configmap/list", configMapController.GetConfigMapList)
	// configmap 删除  /configmap/delete?namespace=dev&name=testcm
	group.GET("/configmap/delete", configMapController.DeleteConfigMap)

	//******************secret************************//
	// 创建或更新secret
	group.POST("/secret/create_update", secretController.CreateOrUpdateSecret)
	// secret详情
	group.GET("/secret/detail", secretController.GetSecretDetail)
	// secret  列表
	group.GET("/secret/list", secretController.GetSecretList)
	// secret 删除
	group.GET("/secret/delete", secretController.DeleteSecret)
}
