/*
*
User: cr-mao
Date: 2024/2/24 21:57
Email: crmao@qq.com
Desc: configmap_controller.go
*/
package controllers

import (
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/configmap_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services"
	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"
	"github.com/gin-gonic/gin"
)

type ConfigMapController struct{}

// 创建或修改configmap
func (c *ConfigMapController) CreateOrUpdateConfigMap(ctx *gin.Context) {
	var configMapReq *configmap_request.ConfigMap
	err := ctx.ShouldBind(&configMapReq)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, err.Error())
		return
	}
	err = services.ConfigMapService.CreateOrUpdateConfigMap(ctx, configMapReq)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}

// 查询configmap详情
func (c *ConfigMapController) GetConfigMapDetail(ctx *gin.Context) {
	name := ctx.Query("name")
	namespace := ctx.Query("namespace")
	detail, err := services.ConfigMapService.GetConfigMapDetail(ctx, namespace, name)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, detail)
}

// 查询configmap列表
func (c *ConfigMapController) GetConfigMapList(ctx *gin.Context) {
	namespace := ctx.Query("namespace")
	res, err := services.ConfigMapService.GetConfigMapList(ctx, namespace)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, res)
}

// 删除
func (c *ConfigMapController) DeleteConfigMap(ctx *gin.Context) {
	err := services.ConfigMapService.DeleteConfigMap(ctx, ctx.Query("namespace"), ctx.Query("name"))
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}
