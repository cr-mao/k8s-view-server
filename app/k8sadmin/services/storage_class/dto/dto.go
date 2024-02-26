/*
*
User: cr-mao
Date: 2024/2/26 13:51
Email: crmao@qq.com
Desc: dto.go
*/
package dto

import (
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
)

type StorageClass struct {
	Name string `json:"name"`
	//Namespace string             `json:"namespace"`
	Labels []global.ListMapItem `json:"labels"`
	//制备器
	Provisioner string `json:"provisioner"`
	//卷绑定参数配置
	MountOptions []string `json:"mountOptions"`
	//制备器入参
	Parameters []global.ListMapItem `json:"parameters"`
	//卷回收策略
	ReclaimPolicy corev1.PersistentVolumeReclaimPolicy `json:"reclaimPolicy"`
	//是否允许扩充
	AllowVolumeExpansion bool `json:"allowVolumeExpansion"`
	//卷绑定模式
	VolumeBindingMode storagev1.VolumeBindingMode `json:"volumeBindingMode"`
	Age               int64                       `json:"age"`
}
