/*
*
User: cr-mao
Date: 2024/2/26 13:09
Email: crmao@qq.com
Desc: dto.go
*/
package dto

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
)

type PersistentVolumeClaim struct {
	Name             string                              `json:"name"`
	Namespace        string                              `json:"namespace"`
	Labels           []global.ListMapItem                `json:"labels"`
	AccessModes      []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	Capacity         int32                               `json:"capacity"`
	Selector         []global.ListMapItem                `json:"selector"`
	StorageClassName string                              `json:"storageClassName"`
	Age              int64                               `json:"age"`
	Volume           string                              `json:"volume"`
	//pvc 状态
	Status corev1.PersistentVolumeClaimPhase `json:"status"`
}
