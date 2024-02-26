/*
*
User: cr-mao
Date: 2024/2/26 13:51
Email: crmao@qq.com
Desc: storage_class_service.go
*/
package storage_class

import (
	"context"
	"fmt"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/storage_class_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/storage_class/dto"
	"github.com/cr-mao/k8s-view-server/infra/conf"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type StorageClassService struct{}

func (s *StorageClassService) GetSCList(ctx context.Context) ([]dto.StorageClass, error) {
	list, err := global.KubeConfigSet.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	scResList := make([]dto.StorageClass, 0)
	for _, item := range list.Items {
		var allowVolumeExpansion bool
		if item.AllowVolumeExpansion != nil {
			allowVolumeExpansion = *item.AllowVolumeExpansion
		}
		mountOptions := make([]string, 0)
		if item.MountOptions != nil {
			mountOptions = item.MountOptions
		}
		var reclaimPolicy corev1.PersistentVolumeReclaimPolicy
		if item.ReclaimPolicy != nil {
			reclaimPolicy = *item.ReclaimPolicy
		}
		scResItem := dto.StorageClass{
			Name:                 item.Name,
			Labels:               global.ToList(item.Labels),
			Provisioner:          item.Provisioner,
			MountOptions:         mountOptions,
			Parameters:           global.ToList(item.Parameters),
			ReclaimPolicy:        reclaimPolicy,
			AllowVolumeExpansion: allowVolumeExpansion,
			Age:                  item.CreationTimestamp.UnixMilli(),
			VolumeBindingMode:    *item.VolumeBindingMode,
		}
		scResList = append(scResList, scResItem)
	}
	return scResList, err
}

func (s *StorageClassService) DeleteSC(ctx context.Context, name string) error {
	return global.KubeConfigSet.StorageV1().StorageClasses().
		Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (s *StorageClassService) CreateSC(ctx context.Context, scReq *storage_class_request.StorageClass) error {
	//判断Provisioner是否在系统支持
	provisionerList := strings.Split(conf.GetString("k8s.provisioner"), ",")
	var flag bool
	for _, val := range provisionerList {
		if scReq.Provisioner == val {
			flag = true
			break
		}
	}
	if !flag {
		err := fmt.Errorf("%s 当前K8S未支持！ ", scReq.Provisioner)
		return err
	}
	sc := storagev1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:   scReq.Name,
			Labels: global.ToMap(scReq.Labels),
		},
		Provisioner:          scReq.Provisioner,
		MountOptions:         scReq.MountOptions,
		VolumeBindingMode:    &scReq.VolumeBindingMode,
		ReclaimPolicy:        &scReq.ReclaimPolicy,
		AllowVolumeExpansion: &scReq.AllowVolumeExpansion,
		Parameters:           global.ToMap(scReq.Parameters),
	}
	_, err := global.KubeConfigSet.StorageV1().StorageClasses().
		Create(ctx, &sc, metav1.CreateOptions{})
	return err
}
