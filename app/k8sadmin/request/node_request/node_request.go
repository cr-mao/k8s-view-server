/*
*
User: cr-mao
Date: 2024/2/24 20:59
Email: crmao@qq.com
Desc: node_request.go
*/
package node_request

import (
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	corev1 "k8s.io/api/core/v1"
)

type UpdatedLabel struct {
	Name   string               `json:"name"`
	Labels []global.ListMapItem `json:"labels"`
}

type UpdatedTaint struct {
	Name   string         `json:"name"`
	Taints []corev1.Taint `json:"taints"`
}
