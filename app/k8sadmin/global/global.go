/*
*
User: cr-mao
Date: 2024/2/22 17:05
Email: crmao@qq.com
Desc: global.go
*/
package global

import "k8s.io/client-go/kubernetes"

var KubeConfigSet *kubernetes.Clientset

type ListMapItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
