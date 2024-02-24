/*
*
User: cr-mao
Date: 2024/2/24 22:06
Email: crmao@qq.com
Desc: configmap_request.go
*/
package configmap_request

import "github.com/cr-mao/k8s-view-server/app/k8sadmin/global"

type ConfigMap struct {
	Name      string               `json:"name"`
	Namespace string               `json:"namespace"`
	Labels    []global.ListMapItem `json:"labels"`
	Data      []global.ListMapItem `json:"data"`
}
