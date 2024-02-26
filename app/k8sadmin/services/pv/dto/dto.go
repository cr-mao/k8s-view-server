/*
*
User: cr-mao
Date: 2024/2/26 12:44
Email: crmao@qq.com
Desc: dto.go
*/
package dto

import (
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	corev1 "k8s.io/api/core/v1"
)

type PersistentVolume struct {
	Name string `json:"name"`
	//pv容量
	Capacity int32 `json:"capacity"`
	//ns 不必传
	//Namespace string             `json:"namespace"`
	Labels []global.ListMapItem `json:"labels"`
	//数据读写权限
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	//pv回收策略
	ReClaimPolicy corev1.PersistentVolumeReclaimPolicy `json:"reClaimPolicy"`
	//todo 待完善
	Status corev1.PersistentVolumePhase `json:"status"`
	//被具备某个pvc绑定
	Claim string `json:"claim"`
	//创建时间
	Age int64 `json:"age"`
	//状况描述
	Reason string `json:"reason"`
	//sc 名称
	StorageClassName string `json:"storageClassName"`
}
