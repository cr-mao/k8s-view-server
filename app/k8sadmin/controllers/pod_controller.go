/*
*
User: cr-mao
Date: 2024/2/22 14:41
Email: crmao@qq.com
Desc: pod_controller.go
*/
package controllers

import (
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/pod_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services"
	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"
	"github.com/gin-gonic/gin"
)

type PodController struct{}

// pod 创建更新
func (c *PodController) CreateOrUpdate(ctx *gin.Context) {
	var podReq *pod_request.Pod
	if err := ctx.ShouldBind(&podReq); err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, err.Error())
		return
	}
	//校验必填项
	if err := pod_request.Validate(podReq); err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, err.Error())
		return
	}
	msg, err := services.PodService.CreateOrUpdatePod(ctx, podReq)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, msg)
}

// 删除pod
func (c *PodController) DeletePod(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	err := services.PodService.DeletePod(ctx, namespace, name)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}
