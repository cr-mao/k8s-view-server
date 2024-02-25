/*
*
User: cr-mao
Date: 2024/2/25 09:53
Email: crmao@qq.com
Desc: secret_request.go
*/
package secret_request

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
)

type Secret struct {
	Name      string               `json:"name"`
	Namespace string               `json:"namespace"`
	Type      corev1.SecretType    `json:"type"`
	Labels    []global.ListMapItem `json:"labels"`
	Data      []global.ListMapItem `json:"data"`
}
