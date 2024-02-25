/*
*
User: cr-mao
Date: 2024/2/24 19:17
Email: crmao@qq.com
Desc: dto.go
*/
package dto

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
)

type Node struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Age        int64  `json:"age"`
	InternalIp string `json:"internalIp"`
	ExternalIp string `json:"externalIp"`
	//kubelet 版本
	Version       string `json:"version"`
	OsImage       string `json:"osImage"`
	KernelVersion string `json:"kernelVersion"`
	//容器运行时
	ContainerRuntime string               `json:"containerRuntime"`
	Labels           []global.ListMapItem `json:"labels"`
	Taints           []corev1.Taint       `json:"taints"`
}
