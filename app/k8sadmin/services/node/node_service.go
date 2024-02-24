/*
*
User: cr-mao
Date: 2024/2/24 19:16
Email: crmao@qq.com
Desc: node_service.go
*/
package node

import (
	"context"
	"encoding/json"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/node_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/node/dto"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type NodeService struct{}

// node详情
func (*NodeService) GetNodeDetail(ctx context.Context, nodeName string) (*dto.Node, error) {
	nodeK8s, err := global.KubeConfigSet.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	detail := getNodeDetail(nodeK8s)
	return &detail, err
}

// node 列表
func (*NodeService) GetNodeList(ctx context.Context) ([]dto.Node, error) {
	list, err := global.KubeConfigSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	nodeResList := make([]dto.Node, 0)
	for _, item := range list.Items {
		nodeRes := getNodeDetail(&item)
		nodeResList = append(nodeResList, nodeRes)
	}
	return nodeResList, err
}

// 污点更新
func (*NodeService) UpdateNodeTaint(ctx context.Context, updatedTaint *node_request.UpdatedTaint) error {
	patchData := map[string]any{
		"spec": map[string]any{
			"taints": updatedTaint.Taints,
		},
	}
	patchDataBytes, _ := json.Marshal(&patchData)
	_, err := global.KubeConfigSet.CoreV1().Nodes().Patch(
		ctx,
		updatedTaint.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	return err
}

// 标签更新
func (*NodeService) UpdateNodeLabel(ctx context.Context, updatedLabel *node_request.UpdatedLabel) error {
	labelsMap := make(map[string]string, 0)
	for _, label := range updatedLabel.Labels {
		labelsMap[label.Key] = label.Value
	}
	// 替换的方式
	labelsMap["$patch"] = "replace"
	patchData := map[string]any{
		"metadata": map[string]any{
			"labels": labelsMap,
		},
	}
	patchDataBytes, _ := json.Marshal(&patchData)
	_, err := global.KubeConfigSet.CoreV1().Nodes().Patch(
		ctx,
		updatedLabel.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	return err
}

func getNodeStatus(nodeConditions []corev1.NodeCondition) string {
	nodeStatus := "NotReady"
	for _, condition := range nodeConditions {
		if condition.Type == "Ready" && condition.Status == "True" {
			nodeStatus = "Ready"
			break
		}
	}
	return nodeStatus
}

func getNodeIp(addresses []corev1.NodeAddress, addressType corev1.NodeAddressType) string {
	for _, item := range addresses {
		if item.Type == addressType {
			return item.Address
		}
	}
	return "<none>"
}

func mapToList(m map[string]string) []global.ListMapItem {
	res := make([]global.ListMapItem, 0)
	for k, v := range m {
		res = append(res, global.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return res
}

func getNodeDetail(nodeK8s *corev1.Node) dto.Node {
	nodeRes := getNodeResItem(nodeK8s)
	//计算label 和 taint
	nodeRes.Taints = nodeK8s.Spec.Taints
	nodeRes.Labels = mapToList(nodeK8s.Labels)
	return nodeRes
}

func getNodeResItem(nodeK8s *corev1.Node) dto.Node {
	nodeInfo := nodeK8s.Status.NodeInfo
	return dto.Node{
		Name:             nodeK8s.Name,
		Status:           getNodeStatus(nodeK8s.Status.Conditions),
		Age:              nodeK8s.CreationTimestamp.Unix(),
		InternalIp:       getNodeIp(nodeK8s.Status.Addresses, corev1.NodeInternalIP),
		ExternalIp:       getNodeIp(nodeK8s.Status.Addresses, corev1.NodeExternalIP),
		OsImage:          nodeInfo.OSImage,
		Version:          nodeInfo.KubeletVersion,
		KernelVersion:    nodeInfo.KernelVersion,
		ContainerRuntime: nodeInfo.ContainerRuntimeVersion,
	}
}
