/*
*
User: cr-mao
Date: 2024/2/24 18:38
Email: crmao@qq.com
Desc: node_controller.go
*/
package controllers

import (
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/node_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services"
	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"
	"github.com/gin-gonic/gin"
)

type NodeController struct{}

// node 详情
func (c *NodeController) GetNodeDetail(ctx *gin.Context) {
	name := ctx.Query("node_name")
	if name == "" {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, "node 名不能为空")
	}
	res, err := services.NodeService.GetNodeDetail(ctx, name)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, res)
}

// node 列表
func (c *NodeController) GetNodeList(ctx *gin.Context) {
	res, err := services.NodeService.GetNodeList(ctx)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, res)
}

// 更新标签
func (c *NodeController) UpdateNodeLabel(ctx *gin.Context) {
	var updatedLabel *node_request.UpdatedLabel
	err := ctx.ShouldBind(&updatedLabel)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, "参数错误")
		return
	}
	err = services.NodeService.UpdateNodeLabel(ctx, updatedLabel)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}

// 污点设置更新
func (*NodeController) UpdateNodeTaint(ctx *gin.Context) {
	var updatedTaint *node_request.UpdatedTaint
	err := ctx.ShouldBind(&updatedTaint)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, "参数错误")
		return
	}
	err = services.NodeService.UpdateNodeTaint(ctx, updatedTaint)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}
