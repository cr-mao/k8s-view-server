/*
*
User: cr-mao
Date: 2024/2/24 22:09
Email: crmao@qq.com
Desc: configmap_service.go
*/
package configmap

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/configmap_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/configmap/dto"
)

type ConfigMapService struct{}

//kubectl get configmap <configmap-name> -o yaml > configmap.yaml

// 创建更新configmap
func (s *ConfigMapService) CreateOrUpdateConfigMap(ctx context.Context, configMapReq *configmap_request.ConfigMap) error {
	// 将 request 转为 k8s 结构
	configMap := cmReq2K8sConvert(configMapReq)
	//判断是否存在
	_, errSearch := global.KubeConfigSet.CoreV1().ConfigMaps(configMapReq.Namespace).Get(ctx, configMapReq.Name, metav1.GetOptions{})
	if errSearch == nil {
		_, err := global.KubeConfigSet.CoreV1().ConfigMaps(configMapReq.Namespace).Update(ctx, configMap, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	} else {
		_, err := global.KubeConfigSet.CoreV1().ConfigMaps(configMapReq.Namespace).Create(ctx, configMap, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

// 详情
func (s *ConfigMapService) GetConfigMapDetail(ctx context.Context, namespace, name string) (*dto.ConfigMap, error) {
	configMapK8s, err := global.KubeConfigSet.CoreV1().ConfigMaps(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	cm := geCmReqDetail(configMapK8s)
	return cm, nil
}

// configmap 列表
func (s *ConfigMapService) GetConfigMapList(ctx context.Context, namespace string) ([]*dto.ConfigMap, error) {
	//1 从k8s查询
	list, err := global.KubeConfigSet.CoreV1().ConfigMaps(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	//2 转换为res(filter)
	configMapList := make([]*dto.ConfigMap, 0)
	for _, item := range list.Items {
		configMapList = append(configMapList, geCmReqItem(&item))
	}
	return configMapList, nil
}

// 删除configmap
func (*ConfigMapService) DeleteConfigMap(ctx context.Context, ns string, name string) error {
	return global.KubeConfigSet.CoreV1().ConfigMaps(ns).Delete(ctx, name, metav1.DeleteOptions{})
}

func geCmReqDetail(configMap *corev1.ConfigMap) *dto.ConfigMap {
	detail := geCmReqItem(configMap)
	detail.Labels = global.ToList(configMap.Labels)
	detail.Data = global.ToList(configMap.Data)
	return detail
}

func geCmReqItem(configMap *corev1.ConfigMap) *dto.ConfigMap {
	return &dto.ConfigMap{
		Name:      configMap.Name,
		Namespace: configMap.Namespace,
		DataNum:   len(configMap.Data),
		Age:       configMap.CreationTimestamp.Unix(),
	}
}

func cmReq2K8sConvert(configMapReq *configmap_request.ConfigMap) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapReq.Name,
			Namespace: configMapReq.Namespace,
			Labels:    global.ToMap(configMapReq.Labels),
		},
		Data: global.ToMap(configMapReq.Data),
	}
}
