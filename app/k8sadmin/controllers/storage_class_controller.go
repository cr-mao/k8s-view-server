/*
*
User: cr-mao
Date: 2024/2/26 13:48
Email: crmao@qq.com
Desc: storage_class_controller.go
*/
package controllers

import (
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/storage_class_request"
	"github.com/gin-gonic/gin"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services"
	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"
)

type StorageClassController struct{}

// StorageClass列表
func (c *StorageClassController) GetSCList(ctx *gin.Context) {
	list, err := services.StorageClassService.GetSCList(ctx)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, list)
}

// StorageClass创建
func (c *StorageClassController) CreateSC(ctx *gin.Context) {
	var scReq *storage_class_request.StorageClass
	if err := ctx.ShouldBind(&scReq); err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, err.Error())
		return
	}
	err := services.StorageClassService.CreateSC(ctx, scReq)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}

// StorageClass删除
func (c *StorageClassController) DeleteSC(ctx *gin.Context) {
	name := ctx.Query("name")
	err := services.StorageClassService.DeleteSC(ctx, name)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}
