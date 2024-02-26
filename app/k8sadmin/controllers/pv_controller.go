/*
*
User: cr-mao
Date: 2024/2/26 12:33
Email: crmao@qq.com
Desc: pv_controller.go
*/
package controllers

import (
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/pv_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services"
	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"
	"github.com/gin-gonic/gin"
)

type PvController struct{}

// 创建pv
func (c *PvController) CreatePV(ctx *gin.Context) {
	var pvReq *pv_request.PersistentVolume
	if err := ctx.ShouldBind(&pvReq); err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, err.Error())
		return
	}
	err := services.PvService.CreatePV(ctx, pvReq)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}

// pv list
func (c *PvController) GetPVList(ctx *gin.Context) {
	list, err := services.PvService.GetPvList(ctx)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, list)
}

// 删除pv
func (c *PvController) DeletePV(ctx *gin.Context) {
	if ctx.Query("name") == "" {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, "name不能为空")
		return
	}
	err := services.PvService.DeletePV(ctx, ctx.Query("name"))
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}


