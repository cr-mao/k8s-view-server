/*
*
User: cr-mao
Date: 2024/2/26 13:00
Email: crmao@qq.com
Desc: pvc_controller.go
*/
package controllers

import (
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/pvc_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services"
	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"
	"github.com/gin-gonic/gin"
)

type PvcController struct{}

// 创建pvc
func (c *PvcController) CreatePVC(ctx *gin.Context) {
	var pvcReq *pvc_request.PersistentVolumeClaim
	if err := ctx.ShouldBind(&pvcReq); err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, err.Error())
		return
	}
	err := services.PvcService.CreatePVC(ctx, pvcReq)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}

// pvc list
func (c *PvcController) GetPVCList(ctx *gin.Context) {
	if ctx.Query("namespace") == "" {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, "namespace不能为空")
		return
	}
	list, err := services.PvcService.GetPVCList(ctx, ctx.Query("namespace"))
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, list)
}

// 删除pvc
func (c *PvcController) DeletePVC(ctx *gin.Context) {
	namespace := ctx.Query("namespace")
	name := ctx.Query("name")

	if ctx.Query("name") == "" || ctx.Query("namespace") == "" {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, "name、namespace不能为空")
		return
	}
	err := services.PvcService.DeletePVC(ctx, namespace, name)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}
