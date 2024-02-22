/*
*
User: cr-mao
Date: 2024/2/22 14:02
Email: crmao@qq.com
Desc: namespace_service.go
*/
package namespace

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/namespace/dto"
)

type NameSpaceService struct{}

// 获得命名空间列表
func (s *NameSpaceService) GetNameSpaceList(ctx context.Context) ([]dto.Namespace, error) {
	list, err := global.KubeConfigSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	namespaceList := make([]dto.Namespace, 0)
	for _, item := range list.Items {
		namespaceList = append(namespaceList, dto.Namespace{
			Name:      item.Name,
			CreatedAt: item.CreationTimestamp.Unix(),
			Status:    string(item.Status.Phase),
		})
	}
	return namespaceList, nil
}
