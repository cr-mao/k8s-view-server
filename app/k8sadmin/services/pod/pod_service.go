/*
*
User: cr-mao
Date: 2024/2/22 16:23
Email: crmao@qq.com
Desc: pod_service.go
*/
package pod

import (
	"context"
	"fmt"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/convert"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/pod_request"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodService struct{}

// 创建更新pod
func (*PodService) CreateOrUpdatePod(ctx context.Context, podReq *pod_request.Pod) (msg string, err error) {
	//[no]update [no]patch [yes]delete+create
	k8sPod := convert.PodReq2K8sConvert.PodReq2K8s(podReq)
	podApi := global.KubeConfigSet.CoreV1().Pods(k8sPod.Namespace)
	createdPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
	if err != nil {
		errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]创建失败，detail：%s", k8sPod.Namespace, k8sPod.Name, err.Error())
		return errMsg, err
	}
	successMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]创建成功", createdPod.Namespace, createdPod.Name)
	return successMsg, err
}

// 删除pod
func (*PodService) DeletePod(ctx context.Context, namespace string, name string) error {
	// https://dandelioncloud.cn/article/details/1507597894772400129
	background := metav1.DeletePropagationBackground
	var gracePeriodSeconds int64 = 0
	return global.KubeConfigSet.CoreV1().Pods(namespace).Delete(ctx, name,
		metav1.DeleteOptions{
			GracePeriodSeconds: &gracePeriodSeconds,
			PropagationPolicy:  &background,
		})
}
