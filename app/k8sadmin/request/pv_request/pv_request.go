/*
*
User: cr-mao
Date: 2024/2/26 12:31
Email: crmao@qq.com
Desc: pv_request.go
*/
package pv_request

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
)

type NfsVolumeSource struct {
	NfsPath     string `json:"nfsPath"`
	NfsServer   string `json:"nfsServer"`
	NfsReadOnly bool   `json:"nfsReadOnly"`
}
type VolumeSource struct {
	Type            string          `json:"type"`
	NfsVolumeSource NfsVolumeSource `json:"nfsVolumeSource"`
}
type PersistentVolume struct {
	Name string `json:"name"`
	//ns 不必传
	//Namespace string             `json:"namespace"`
	Labels []global.ListMapItem `json:"labels"`
	//pv容量
	Capacity int32 `json:"capacity"`
	//数据读写权限
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	//pv回收策略
	ReClaimPolicy corev1.PersistentVolumeReclaimPolicy `json:"reClaimPolicy"`
	VolumeSource  VolumeSource                         `json:"volumeSource"`
}
