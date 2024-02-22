/*
*
User: cr-mao
Date: 2024/2/22 13:59
Email: crmao@qq.com
Desc: namespace_controller.go
*/
package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services/namespace"
	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"
)

type NamespaceController struct{}

// 命名空间列表
func (cc *NamespaceController) GetNamespaceList(ctx *gin.Context) {
	service := &namespace.NameSpaceService{}
	res, err := service.GetNameSpaceList(ctx)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, res)
}
