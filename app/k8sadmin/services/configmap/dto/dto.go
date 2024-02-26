/*
*
User: cr-mao
Date: 2024/2/24 22:57
Email: crmao@qq.com
Desc: dto.go
*/
package dto

import "github.com/cr-mao/k8s-view-server/app/k8sadmin/global"

type ConfigMap struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	DataNum   int    `json:"dataNum"`
	Age       int64  `json:"age"`
	//查询configmap详情信息
	Data   []global.ListMapItem `json:"data"`
	Labels []global.ListMapItem `json:"labels"`
}
