/*
*
User: cr-mao
Date: 2024/2/26 13:04
Email: crmao@qq.com
Desc: pvc_service.go
*/
package pvc

import (
	"context"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/pvc_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/pvc/dto"
)

type PvcService struct{}

func (s *PvcService) GetPVCList(ctx context.Context, namespace string) ([]dto.PersistentVolumeClaim, error) {
	pvcResList := make([]dto.PersistentVolumeClaim, 0)
	list, err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(namespace).
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range list.Items {
		//item -> response
		matchLabels := make([]global.ListMapItem, 0)
		if item.Spec.Selector != nil {
			matchLabels = global.ToList(item.Spec.Selector.MatchLabels)
		}
		storageClassName := ""
		if item.Spec.StorageClassName != nil {
			storageClassName = *item.Spec.StorageClassName
		}
		pvcResItem := dto.PersistentVolumeClaim{
			Name:      item.Name,
			Namespace: item.Namespace,
			Status:    item.Status.Phase,
			//转换为Mi
			Capacity:         int32(item.Spec.Resources.Requests.Storage().Value() / (1024 * 1024)),
			AccessModes:      item.Spec.AccessModes,
			StorageClassName: storageClassName,
			Age:              item.CreationTimestamp.UnixMilli(),
			Volume:           item.Spec.VolumeName,
			Labels:           global.ToList(item.Labels),
			Selector:         matchLabels,
		}
		pvcResList = append(pvcResList, pvcResItem)
	}
	return pvcResList, err
}

// 删除pvc
func (s *PvcService) DeletePVC(ctx context.Context, namespace string, name string) error {
	err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(namespace).
		Delete(ctx, name, metav1.DeleteOptions{})
	return err
}

// 创建pvc
func (s *PvcService) CreatePVC(ctx context.Context, pvcReq *pvc_request.PersistentVolumeClaim) error {
	pvc := corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pvcReq.Name,
			Namespace: pvcReq.Namespace,
			Labels:    global.ToMap(pvcReq.Labels),
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: global.ToMap(pvcReq.Selector),
			},
			AccessModes: pvcReq.AccessModes,
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(strconv.Itoa(int(pvcReq.Capacity)) + "Mi"),
				},
			},
			StorageClassName: &pvcReq.StorageClassName,
		},
	}
	if pvc.Spec.StorageClassName != nil {
		pvc.Spec.Selector = nil
	}
	_, err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(pvc.Namespace).
		Create(ctx, &pvc, metav1.CreateOptions{})
	return err
}
