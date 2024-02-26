/*
*
User: cr-mao
Date: 2024/2/26 12:35
Email: crmao@qq.com
Desc: pv_service.go
*/
package pv

import (
	"context"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/pv/dto"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/pv_request"
	"github.com/cr-mao/k8s-view-server/infra/errors"
)

type PvService struct{}

// 创建pv
func (s *PvService) CreatePV(ctx context.Context, pvReq *pv_request.PersistentVolume) error {
	//参数转换
	var volumeSource corev1.PersistentVolumeSource
	switch pvReq.VolumeSource.Type {
	case "nfs":
		volumeSource.NFS = &corev1.NFSVolumeSource{
			Server:   pvReq.VolumeSource.NfsVolumeSource.NfsServer,
			Path:     pvReq.VolumeSource.NfsVolumeSource.NfsPath,
			ReadOnly: pvReq.VolumeSource.NfsVolumeSource.NfsReadOnly,
		}
	default:
		return errors.New("不支持的存储卷类型！")
	}
	pv := corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:   pvReq.Name,
			Labels: global.ToMap(pvReq.Labels),
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceStorage: resource.MustParse(strconv.Itoa(int(pvReq.Capacity)) + "Mi"),
			},
			AccessModes:                   pvReq.AccessModes,
			PersistentVolumeReclaimPolicy: pvReq.ReClaimPolicy,
			PersistentVolumeSource:        volumeSource,
		},
	}
	_, err := global.KubeConfigSet.CoreV1().PersistentVolumes().Create(ctx, &pv, metav1.CreateOptions{})
	return err
}

// pv 列表
func (s *PvService) GetPvList(ctx context.Context) ([]dto.PersistentVolume, error) {
	pvList, err := global.KubeConfigSet.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	pvResList := make([]dto.PersistentVolume, 0)
	for _, item := range pvList.Items {
		//k8s -> response
		claim := ""
		if item.Spec.ClaimRef != nil {
			claim = item.Spec.ClaimRef.Name
		}
		pvRes := dto.PersistentVolume{
			Name:          item.Name,
			Labels:        global.ToList(item.Labels),
			Capacity:      int32(item.Spec.Capacity.Storage().Value() / (1024 * 1024)),
			AccessModes:   item.Spec.AccessModes,
			ReClaimPolicy: item.Spec.PersistentVolumeReclaimPolicy,
			Status:        item.Status.Phase,
			Claim:         claim,
			// 当pv是通过sc创建时 就会有该字段
			StorageClassName: item.Spec.StorageClassName,
			Reason:           item.Status.Reason,
			Age:              item.CreationTimestamp.UnixMilli(),
		}
		pvResList = append(pvResList, pvRes)
	}
	return pvResList, err
}

func (s *PvService) DeletePV(ctx context.Context, name string) error {
	err := global.KubeConfigSet.CoreV1().PersistentVolumes().Delete(ctx, name, metav1.DeleteOptions{})
	return err
}
