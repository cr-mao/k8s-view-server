/*
*
User: cr-mao
Date: 2024/2/25 09:55
Email: crmao@qq.com
Desc: secret_service.go
*/
package secret

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/secret_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/secret/dto"
)

type SecretService struct{}

// 创建更新secret
func (s *SecretService) CreateOrUpdateSecret(ctx context.Context, secretReq *secret_request.Secret) (err error) {
	_, err = global.KubeConfigSet.CoreV1().Secrets(secretReq.Namespace).Get(ctx, secretReq.Name, metav1.GetOptions{})
	k8sSecret := secret2K8sSecret(secretReq)
	// 更新
	if err == nil {
		_, err = global.KubeConfigSet.CoreV1().Secrets(secretReq.Namespace).Update(ctx, k8sSecret, metav1.UpdateOptions{})
	} else {
		// 创建
		_, err = global.KubeConfigSet.CoreV1().Secrets(secretReq.Namespace).Create(ctx, k8sSecret, metav1.CreateOptions{})
	}
	return
}

// 删除secret
func (s *SecretService) DeleteSecret(ctx context.Context, namespace string, name string) error {
	return global.KubeConfigSet.CoreV1().Secrets(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

// secret 详情
func (SecretService) GetSecretDetail(ctx context.Context, namespace string, name string) (*dto.Secret, error) {
	secretK8s, err := global.KubeConfigSet.CoreV1().Secrets(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	secretRes := secretK8s2ResDetailConvert(secretK8s)
	return secretRes, err
}

// secret 列表
func (SecretService) GetSecretList(ctx context.Context, namespace string) ([]dto.Secret, error) {
	list, err := global.KubeConfigSet.CoreV1().Secrets(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	secretResList := make([]dto.Secret, 0)
	for _, item := range list.Items {
		secretRes := secretK8s2ResItemConvert(&item)
		secretResList = append(secretResList, secretRes)
	}
	return secretResList, err
}

func secretK8s2ResItemConvert(secret *corev1.Secret) dto.Secret {
	return dto.Secret{
		Name:      secret.Name,
		Namespace: secret.Namespace,
		Type:      secret.Type,
		DataNum:   len(secret.Data),
		Age:       secret.CreationTimestamp.Unix(),
	}
}

func secretK8s2ResDetailConvert(secret *corev1.Secret) *dto.Secret {
	return &dto.Secret{
		Name:      secret.Name,
		Namespace: secret.Namespace,
		Type:      secret.Type,
		DataNum:   len(secret.Data),
		Age:       secret.CreationTimestamp.Unix(),
		Data:      global.ToListWithMapByte(secret.Data),
		Labels:    global.ToList(secret.Labels),
	}
}

func secret2K8sSecret(secretReq *secret_request.Secret) *corev1.Secret {
	return &corev1.Secret{
		Type: secretReq.Type,
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretReq.Name,
			Namespace: secretReq.Namespace,
			Labels:    global.ToMap(secretReq.Labels),
		},
		StringData: global.ToMap(secretReq.Data),
	}
}
