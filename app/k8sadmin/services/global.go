/*
*
User: cr-mao
Date: 2024/2/22 17:03
Email: crmao@qq.com
Desc: global.go
*/
package services

import (
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/configmap"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/namespace"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/node"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/pod"
)

var NamespaceService = &namespace.NameSpaceService{}
var PodService = &pod.PodService{}
var NodeService = &node.NodeService{}
var ConfigMapService = &configmap.ConfigMapService{}
